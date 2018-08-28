package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// Issuer issues out problems
type Issuer interface {
	createIssue(service, methodCall string, timeMs int32, otherCalls string)
	getSentCount() int
}

// Server the main server type
type Server struct {
	*goserver.GoServer
	logDirectory  string
	stats         []*pb.Stats
	logs          []*pb.MessageLog
	issuer        Issuer
	LastSlowCheck time.Time
	RunTimeLock   *sync.Mutex
}

//ClearStats clears out the stats
func (s *Server) ClearStats(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	s.stats = make([]*pb.Stats, 0)
	return &pb.Empty{}, nil
}

func (s *Server) emailSlowFunction() {
	s.LastSlowCheck = time.Now()
	for _, st := range s.stats {
		if st.GetMeanRunTime() > 500 {
			//Build up a super string
			super := ""
			for _, sti := range s.stats {
				super += fmt.Sprintf("Also %v/%v -> %v\n", sti.GetBinary(), sti.GetName(), sti.GetMeanRunTime())
			}

			s.issuer.createIssue(st.GetBinary(), st.GetName(), st.GetMeanRunTime(), super)
			return
		}
	}
}

// InitServer creates a monitoring server
func InitServer() *Server {
	s := &Server{&goserver.GoServer{}, "logs", make([]*pb.Stats, 0), make([]*pb.MessageLog, 0), ProdIssuer{}, time.Now(), &sync.Mutex{}}
	s.issuer = ProdIssuer{Resolver: s.GetIP}
	s.Register = s
	return s
}

// WriteMessageLog Writes out a message log
func (s *Server) WriteMessageLog(ctx context.Context, in *pb.MessageLog) (*pb.LogWriteResponse, error) {
	in.Timestamps = time.Now().Unix()
	s.logs = append(s.logs, in)

	//Keep the # of logs fixed to 500
	if len(s.logs) > 500 {
		s.logs = s.logs[len(s.logs)-500:]
	}

	return &pb.LogWriteResponse{Success: true, Timestamp: in.Timestamps}, nil
}

// ReadMessageLogs Reads and returns the message logs for a given entry
func (s *Server) ReadMessageLogs(ctx context.Context, in *pbr.RegistryEntry) (*pb.MessageLogReadResponse, error) {
	response := &pb.MessageLogReadResponse{Logs: make([]*pb.MessageLog, 0)}
	for _, log := range s.logs {
		if log != nil && log.Entry != nil && (log.Entry.Name == in.GetName() || in.GetName() == "") {
			response.Logs = append(response.Logs, log)
		}
	}
	return response, nil
}

func recompute(st *pb.Stats) {
	sum := int32(0)
	for _, t := range st.RunTimes {
		sum += t
	}
	st.MeanRunTime = sum / int32(len(st.RunTimes))
}

// GetStats gets the stats for a given function call
func (s *Server) GetStats(ctx context.Context, in *pb.FunctionCall) (*pb.StatsList, error) {
	if in.Binary == "" && in.Name == "" {
		for _, st := range s.stats {
			recompute(st)
		}
		return &pb.StatsList{Stats: s.stats}, nil
	}

	for _, st := range s.stats {
		if st.Binary == in.Binary && st.Name == in.Name {
			recompute(st)
			return &pb.StatsList{Stats: []*pb.Stats{st}}, nil
		}
	}

	return nil, errors.New("Unable to find stats")
}

//WriteFunctionCall writes a function call to the monitoring
func (s *Server) WriteFunctionCall(ctx context.Context, in *pb.FunctionCall) (*pb.Empty, error) {
	if in.Time == 0 {
		return &pb.Empty{}, nil
	}

	var st *pb.Stats
	for _, sta := range s.stats {
		if sta.Binary == in.Binary && sta.Name == in.Name {
			st = sta
		}
	}

	if st == nil {
		st = &pb.Stats{Name: in.Name, Binary: in.Binary, NumberOfCalls: 1, MeanRunTime: in.Time, RunTimes: []int32{in.Time}, Slowest: in}
		s.stats = append(s.stats, st)
	} else {
		s.RunTimeLock.Lock()
		st.NumberOfCalls++
		st.MeanRunTime = (st.MeanRunTime*(st.NumberOfCalls-1))/st.NumberOfCalls + in.Time/st.NumberOfCalls
		st.RunTimes = append(st.RunTimes, in.Time)

		if len(st.RunTimes) > numCalls {
			st.RunTimes = st.RunTimes[len(st.RunTimes)-numCalls : len(st.RunTimes)-1]
		}

		if in.Time > st.Slowest.Time {
			st.Slowest = in
		}

		s.RunTimeLock.Unlock()
	}

	return &pb.Empty{}, nil
}
