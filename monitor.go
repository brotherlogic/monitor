package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbgs "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/monitor/proto"
)

const (
	//ConfigKey is where we store the overall config
	ConfigKey = "github.com/brotherlogic/monitor/config"
)

// Server the main server type
type Server struct {
	*goserver.GoServer
	logDirectory  string
	logs          map[string]*logHolder
	logsMutex     *sync.Mutex
	LastSlowCheck time.Time
	RunTimeLock   *sync.Mutex
	reads         int
	writes        int
	writeMap      map[string]int
	writeMutex    *sync.Mutex
	config        *pb.Config
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
	if master {
		return s.load(ctx)
	}
	return nil
}

func (s *Server) save(ctx context.Context) {
	s.KSclient.Save(ctx, ConfigKey, s.config)
}

func (s *Server) load(ctx context.Context) error {
	config := &pb.Config{}
	data, _, err := s.KSclient.Read(ctx, ConfigKey, config)

	if err != nil {
		return err
	}

	s.config = data.(*pb.Config)

	for _, log := range s.config.Logs {
		if log.Entry != nil {
			if _, ok := s.logs[log.Entry.Name]; !ok {
				s.logs[log.Entry.Name] = &logHolder{
					logs:    make([]*pb.MessageLog, 200),
					pointer: 0,
				}
			}

			s.logs[log.Entry.Name].logs[s.logs[log.Entry.Name].pointer] = log
			s.logs[log.Entry.Name].pointer++
			s.logs[log.Entry.Name].pointer %= 200
		}
	}

	return nil
}

// GetState gets the state of the server
func (s Server) GetState() []*pbgs.State {
	s.writeMutex.Lock()
	defer s.writeMutex.Unlock()
	logsLen := 0
	s.logsMutex.Lock()
	for _, val := range s.logs {
		logsLen += len(val.logs)
	}
	s.logsMutex.Unlock()

	maxWrites := 0
	maxWriter := ""
	for key, val := range s.writeMap {
		if val > maxWrites {
			maxWrites = val
			maxWriter = key
		}
	}

	return []*pbgs.State{
		&pbgs.State{Key: "count_reads", Value: int64(s.reads)},
		&pbgs.State{Key: "count_writes", Value: int64(s.writes)},
		&pbgs.State{Key: "max_writer", Text: maxWriter},
		&pbgs.State{Key: "stored_logs", Value: int64(len(s.config.Logs))},
		&pbgs.State{Key: "total_logs", Value: int64(logsLen)},
	}
}

// InitServer creates a monitoring server
func InitServer() *Server {
	s := &Server{
		&goserver.GoServer{},
		"logs",
		make(map[string]*logHolder),
		&sync.Mutex{},
		time.Now(),
		&sync.Mutex{},
		0,
		0,
		make(map[string]int),
		&sync.Mutex{},
		&pb.Config{},
	}
	s.Register = s
	return s
}

func main() {
	var init = flag.Bool("init", false, "Initialise the config")
	flag.Parse()
	s := InitServer()
	s.GoServer.KSclient = *keystoreclient.GetClient(s.DialMaster)
	s.PrepServer()

	// Monitor can't call the monitor
	s.SkipLog = true
	s.SendTrace = false

	err := s.RegisterServerV2("monitor", false, false)
	if err != nil {
		return
	}

	if *init {
		ctx, cancel := utils.BuildContext("monitor", "monitor")
		defer cancel()
		s.config.Logs = append(s.config.Logs, &pb.MessageLog{Message: "Init"})
		s.save(ctx)
	}

	err = s.Serve()
	if err != nil {
		log.Fatalf("Error serving: %v", err)

	}
}
