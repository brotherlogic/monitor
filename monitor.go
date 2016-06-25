package main

import (
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/monitor/monitorproto"
)

// Register registers the rpc end point
func (s *Server) Register(server *grpc.Server) {
	pb.RegisterMonitorServiceServer(server, s)
}

func main() {
	s := InitServer()
	s.RegisterServer("monitor", true)
	s.Serve()
}
