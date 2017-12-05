package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// Issuer issues out problems
type Issuer interface {
	createIssue(service, methodCall string, timeMs int32)
	getSentCount() int
}

// Server the main server type
type Server struct {
	*goserver.GoServer
	heartbeats    []*pb.Heartbeat
	logDirectory  string
	stats         []*pb.Stats
	logs          []*pb.MessageLog
	issuer        Issuer
	LastSlowCheck time.Time
}

func (s *Server) emailSlowFunction() {
	s.LastSlowCheck = time.Now()
	for _, st := range s.stats {
		if st.GetMeanRunTime() > 500 {
			s.WriteMessageLog(context.Background(), &pb.MessageLog{Entry: s.Registry, Message: fmt.Sprintf("Creating issue: %v, %v, %v", st.GetBinary(), st.GetName(), st.GetMeanRunTime())})
			s.issuer.createIssue(st.GetBinary(), st.GetName(), st.GetMeanRunTime())
			return
		}
	}
}

// InitServer creates a monitoring server
func InitServer() *Server {
	s := &Server{&goserver.GoServer{}, make([]*pb.Heartbeat, 0), "logs", make([]*pb.Stats, 0), make([]*pb.MessageLog, 0), ProdIssuer{}, time.Now()}
	s.issuer = ProdIssuer{Resolver: s.GetIP}
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
	in.Timestamps = time.Now().Unix()
	s.logs = append(s.logs, in)
	return &pb.LogWriteResponse{Success: true, Timestamp: in.Timestamps}, nil
}

// ReadMessageLogs Reads and returns the message logs for a given entry
func (s *Server) ReadMessageLogs(ctx context.Context, in *pbr.RegistryEntry) (*pb.MessageLogReadResponse, error) {
	response := &pb.MessageLogReadResponse{}
	for _, log := range s.logs {
		response.Logs = append(response.Logs, log)
	}
	return response, nil
}

func recompute(st *pb.Stats) {
	sum := int32(0)
	for _, t := range st.RunTimes {
		sum += t
	}
	st.MeanRunTime = sum / int32(len(st.RunTimes))
}

// GetStats gets the stats for a given function call
func (s *Server) GetStats(ctx context.Context, in *pb.FunctionCall) (*pb.StatsList, error) {
	if in.Binary == "" && in.Name == "" {
		for _, st := range s.stats {
			recompute(st)
		}
		return &pb.StatsList{Stats: s.stats}, nil
	}

	for _, st := range s.stats {
		if st.Binary == in.Binary && st.Name == in.Name {
			recompute(st)
			return &pb.StatsList{Stats: []*pb.Stats{st}}, nil
		}
	}

	return nil, errors.New("Unable to find stats")
}

//WriteFunctionCall writes a function call to the monitoring
func (s *Server) WriteFunctionCall(ctx context.Context, in *pb.FunctionCall) (*pb.Empty, error) {
	if in.Time == 0 {
		return &pb.Empty{}, nil
	}

	var st *pb.Stats
	for _, sta := range s.stats {
		if sta.Binary == in.Binary && sta.Name == in.Name {
			st = sta
		}
	}

	if st == nil {
		st = &pb.Stats{Name: in.Name, Binary: in.Binary, NumberOfCalls: 1, MeanRunTime: in.Time, RunTimes: []int32{in.Time}}
		s.stats = append(s.stats, st)
	} else {
		st.NumberOfCalls++
		st.MeanRunTime = (st.MeanRunTime*(st.NumberOfCalls-1))/st.NumberOfCalls + in.Time/st.NumberOfCalls
		st.RunTimes = append(st.RunTimes, in.Time)

		if len(st.RunTimes) > numCalls {
			st.RunTimes = st.RunTimes[len(st.RunTimes)-numCalls : len(st.RunTimes)-1]
		}
	}

	return &pb.Empty{}, nil
}

// GetHeartbeats gets the list of per job heartbeats
func (s *Server) GetHeartbeats(ctx context.Context, in *pb.Empty) (*pb.HeartbeatList, error) {
	var mapper map[string]*pb.Heartbeat
	mapper = make(map[string]*pb.Heartbeat)

	for _, heartbeat := range s.heartbeats {
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
