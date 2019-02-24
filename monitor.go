package main

import (
	"fmt"
	"log"
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
	logs          map[string][]*pb.MessageLog
	logsMutex     *sync.Mutex
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

// Shutdown the server
func (s Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s Server) GetState() []*pbgs.State {
	s.writeMutex.Lock()
	defer s.writeMutex.Unlock()
	logsLen := 0
	s.logsMutex.Lock()
	for _, val := range s.logs {
		logsLen += len(val)
	}
	s.logsMutex.Unlock()
	return []*pbgs.State{
		&pbgs.State{Key: "logs", Value: int64(logsLen)},
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
		make(map[string][]*pb.MessageLog),
		&sync.Mutex{},
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
	s.SendTrace = false
	err := s.RegisterServer("monitor", false)
	if err != nil {
		log.Fatalf("Error registering: %v", err)
	}

	err = s.Serve()
	if err != nil {
		log.Fatalf("Error serving: %v", err)

	}
}
