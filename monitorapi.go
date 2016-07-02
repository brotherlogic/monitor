package main

import (
	"github.com/brotherlogic/goserver"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// Server the main server type
type Server struct {
	*goserver.GoServer
	heartbeats   []*pb.Heartbeat
	logDirectory string
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
	s := Server{&goserver.GoServer{}, make([]*pb.Heartbeat, 0), "logs"}
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
	data, _ := proto.Marshal(in)
	ioutil.WriteFile(path, data, 0644)

	return &pb.LogWriteResponse{Success: true, Timestamp: timestamp}, nil
}

// WriteValueLog Writes out a value log
func (s *Server) WriteValueLog(ctx context.Context, in *pb.ValueLog) (*pb.LogWriteResponse, error) {
	path, timestamp := s.getLogPath(in.Entry.Name, in.Entry.Identifier, "value")
	data, _ := proto.Marshal(in)
	ioutil.WriteFile(path, data, 0644)

	return &pb.LogWriteResponse{Success: true, Timestamp: timestamp}, nil
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
