package main

import (
	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"

	"log"
	"time"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// Server the main server type
type Server struct {
	*goserver.GoServer
	heartbeats []*pb.Heartbeat
}

// InitServer creates a monitoring server
func InitServer() Server {
	s := Server{&goserver.GoServer{}, make([]*pb.Heartbeat, 0)}
	s.SetRegisterable(&s)
	return s
}

// ReceiveHeartbeat receives a heartbeat from another server
func (s *Server) ReceiveHeartbeat(ctx context.Context, in *pbr.RegistryEntry) (*pb.Heartbeat, error) {
	heartbeat := &pb.Heartbeat{Entry: in, BeatTime: time.Now().Unix()}
	s.heartbeats = append(s.heartbeats, heartbeat)
	return heartbeat, nil
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
