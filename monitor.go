package main

import (
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/monitor/monitorproto"
)

const (
	//The number of function call details to keep
	numCalls = 100
)

// DoRegister Registers this server
func (s Server) DoRegister(server *grpc.Server) {
	pb.RegisterMonitorServiceServer(server, &s)
}

// ReportHealth alerts if we're not healthy
func (s Server) ReportHealth() bool {
	return true
}

// Mote promotes/demotes this server
func (s Server) Mote(master bool) error {
	return nil
}

func main() {
	s := InitServer()
	s.PrepServer()
	s.GoServer.Killme = true
	s.RegisterServer("monitor", true)
	s.Serve()
}
