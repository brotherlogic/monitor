package main

import (
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/monitor/monitorproto"
)

// DoRegister Registers this server
func (s Server) DoRegister(server *grpc.Server) {
	pb.RegisterMonitorServiceServer(server, &s)
}

// ReportHealth alerts if we're not healthy
func (s Server) ReportHealth() bool {
	return true
}

func main() {
	s := InitServer()
	s.PrepServer()
	s.RegisterServer("monitor", true)
	s.Serve()
}
