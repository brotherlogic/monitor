package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/brotherlogic/goserver"
	pbgs "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// Server the main server type
type Server struct {
	*goserver.GoServer
	logDirectory  string
	logs          []*pb.MessageLog
	LastSlowCheck time.Time
	RunTimeLock   *sync.Mutex
	reads         int
	writes        int
	writeMap      map[string]int
	writeMutex    *sync.Mutex
}

const (
	//The number of function call details to keep
	numCalls = 10000
)

// DoRegister Registers this server
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterMonitorServiceServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s Server) ReportHealth() bool {
	return true
}

// Mote promotes/demotes this server
func (s Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s Server) GetState() []*pbgs.State {
	return []*pbgs.State{
		&pbgs.State{Key: "last_slow", TimeValue: s.LastSlowCheck.Unix()},
		&pbgs.State{Key: "reads", Value: int64(s.reads)},
		&pbgs.State{Key: "writes", Value: int64(s.writes)},
		&pbgs.State{Key: "write_map", Text: fmt.Sprintf("%v", s.writeMap)},
	}
}

// InitServer creates a monitoring server
func InitServer() *Server {
	s := &Server{
		&goserver.GoServer{},
		"logs",
		make([]*pb.MessageLog, 0),
		time.Now(),
		&sync.Mutex{},
		0,
		0,
		make(map[string]int),
		&sync.Mutex{},
	}
	s.Register = s
	return s
}

func main() {
	s := InitServer()
	s.PrepServer()
	s.GoServer.Killme = true
	if s.RegisterServer("monitor", false) {
		err := s.Serve()
		if err != nil {
			fmt.Printf("Error serving: %v", err)
		}
	}
}
