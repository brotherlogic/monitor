package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// Server the main server type
type Server struct {
	*goserver.GoServer
	heartbeats   []*pb.Heartbeat
	logDirectory string
	write        bool
	stats        []*pb.Stats
}

func (s *Server) getLogPath(name string, identifier string, logType string) (string, int64) {
	path := s.logDirectory + "/" + name + "/" + identifier + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}

	timestamp := time.Now().Unix()
	return path + strconv.Itoa(int(timestamp)) + "." + logType, timestamp
}

// InitServer creates a monitoring server
func InitServer() Server {
	s := Server{&goserver.GoServer{}, make([]*pb.Heartbeat, 0), "logs", false, make([]*pb.Stats, 0)}
	s.Register = s
	return s
}

// ReceiveHeartbeat receives a heartbeat from another server
func (s *Server) ReceiveHeartbeat(ctx context.Context, in *pbr.RegistryEntry) (*pb.Heartbeat, error) {
	heartbeat := &pb.Heartbeat{Entry: in, BeatTime: time.Now().Unix()}
	s.heartbeats = append(s.heartbeats, heartbeat)
	return heartbeat, nil
}

// WriteMessageLog Writes out a message log
func (s *Server) WriteMessageLog(ctx context.Context, in *pb.MessageLog) (*pb.LogWriteResponse, error) {
	path, timestamp := s.getLogPath(in.Entry.Name, in.Entry.Identifier, "message")
	in.Timestamps = timestamp
	data, _ := proto.Marshal(in)

	if s.write {
		ioutil.WriteFile(path, data, 0644)
	}

	return &pb.LogWriteResponse{Success: true, Timestamp: timestamp}, nil
}

// ReadMessageLogs Reads and returns the message logs for a given entry
func (s *Server) ReadMessageLogs(ctx context.Context, in *pbr.RegistryEntry) (*pb.MessageLogReadResponse, error) {
	path := s.logDirectory + "/" + in.Name + "/" + in.Identifier + "/"
	files, _ := ioutil.ReadDir(path)
	response := &pb.MessageLogReadResponse{}
	for _, file := range files {
		data, _ := ioutil.ReadFile(path + file.Name())
		logPb := &pb.MessageLog{}
		proto.Unmarshal(data, logPb)
		response.Logs = append(response.Logs, logPb)
	}
	return response, nil
}

// WriteValueLog Writes out a value log
func (s *Server) WriteValueLog(ctx context.Context, in *pb.ValueLog) (*pb.LogWriteResponse, error) {
	path, timestamp := s.getLogPath(in.Entry.Name, in.Entry.Identifier, "value")
	data, _ := proto.Marshal(in)

	if s.write {
		ioutil.WriteFile(path, data, 0644)
	}
	return &pb.LogWriteResponse{Success: true, Timestamp: timestamp}, nil
}

// GetStats gets the stats for a given function call
func (s *Server) GetStats(ctx context.Context, in *pb.FunctionCall) (*pb.StatsList, error) {
	if in.Binary == "" && in.Name == "" {
		return &pb.StatsList{Stats: s.stats}, nil
	}

	for _, st := range s.stats {
		if st.Binary == in.Binary && st.Name == in.Name {
			return &pb.StatsList{Stats: []*pb.Stats{st}}, nil
		}
	}

	return nil, errors.New("Unable to find stats")
}

//WriteFunctionCall writes a function call to the monitoring
func (s *Server) WriteFunctionCall(ctx context.Context, in *pb.FunctionCall) (*pb.Empty, error) {
	log.Printf("WriteFunctionCall %v", in)

	if in.Time == 0 {
		log.Printf("REJECTING CALL AS ZERO TIME")
		return &pb.Empty{}, nil
	}

	var st *pb.Stats
	for _, sta := range s.stats {
		log.Printf("Trying: %v -> %v", sta, in)
		if sta.Binary == in.Binary && sta.Name == in.Name {
			st = sta
		}
	}

	if st == nil {
		st = &pb.Stats{Name: in.Name, Binary: in.Binary, NumberOfCalls: 1, MeanRunTime: in.Time}
		s.stats = append(s.stats, st)
	} else {
		st.NumberOfCalls++
		log.Printf("WORKING WITH %v", st)
		st.MeanRunTime = (st.MeanRunTime*(st.NumberOfCalls-1))/st.NumberOfCalls + in.Time/st.NumberOfCalls
		log.Printf("COMPUTED %v", st.MeanRunTime)
	}

	return &pb.Empty{}, nil
}

// GetHeartbeats gets the list of per job heartbeats
func (s *Server) GetHeartbeats(ctx context.Context, in *pb.Empty) (*pb.HeartbeatList, error) {
	var mapper map[string]*pb.Heartbeat
	mapper = make(map[string]*pb.Heartbeat)

	for _, heartbeat := range s.heartbeats {
		log.Printf("BEAT: %v", heartbeat)
		name := heartbeat.Entry.Identifier + "/" + heartbeat.Entry.Name
		if _, ok := mapper[name]; ok {
			if heartbeat.BeatTime > mapper[name].BeatTime {
				mapper[name] = heartbeat
			}
		} else {
			mapper[name] = heartbeat

		}
	}

	result := &pb.HeartbeatList{}
	for _, beat := range mapper {
		result.Beats = append(result.Beats, beat)
	}

	return result, nil
}
