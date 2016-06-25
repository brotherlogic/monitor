package main

import (
	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"time"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/proto"
)

// Server the main server type
type Server struct {
	goserver.GoServer
	heartbeats []*pb.Heartbeat
}

// InitServer creates a monitoring server
func InitServer() Server {
	s := Server{}
	s.heartbeats = make([]*pb.Heartbeat, 0)
	return s
}

// ReceiveHeartbeat receives a heartbeat from another server
func (s *Server) ReceiveHeartbeat(ctx context.Context, in *pbr.RegistryEntry) (*pb.Heartbeat, error) {
	heartbeat := &pb.Heartbeat{Entry: in, BeatTime: time.Now().Unix()}
	s.heartbeats = append(s.heartbeats, heartbeat)
	return heartbeat, nil
}
