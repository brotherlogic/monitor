package main

import (
	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/proto"
)

// Server the main server type
type Server struct {
	goserver.GoServer
}

// InitServer creates a monitoring server
func InitServer() Server {
	s := Server{}
	return s
}

// ReceiveHeartbeat receives a heartbeat from another server
func (s *Server) ReceiveHeartbeat(ctx context.Context, in *pbr.RegistryEntry) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
