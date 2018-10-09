package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbgh "github.com/brotherlogic/githubcard/proto"
	"github.com/brotherlogic/goserver"
	pbgs "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// ProdIssuer the issuer to use in prod
type ProdIssuer struct {
	Resolver func(string) (string, int)
}

func (p ProdIssuer) createIssue(ctx context.Context, service, methodCall string, timeMs int32, otherCalls string) {
	ip, port := p.Resolver("githubcard")
	if port > 0 {
		conn, _ := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
		defer conn.Close()
		client := pbgh.NewGithubClient(conn)
		client.AddIssue(ctx, &pbgh.Issue{Service: service, Title: "Fix performance", Body: fmt.Sprintf("Fix %v and %v -> %v given %v", service, methodCall, timeMs, otherCalls)})
	}
}

func (p ProdIssuer) getSentCount() int {
	return 0
}

const (
	//The number of function call details to keep
	numCalls = 10000
)

func (s *Server) emailRunner(ctx context.Context) {
	for true {
		s.emailSlowFunction(ctx)
		time.Sleep(time.Hour)
	}
}

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
	}
}

// InitServer creates a monitoring server
func InitServer() *Server {
	s := &Server{
		&goserver.GoServer{},
		"logs",
		make([]*pb.Stats, 0),
		make([]*pb.MessageLog, 0),
		ProdIssuer{},
		time.Now(),
		&sync.Mutex{},
		0,
		0,
	}
	s.issuer = ProdIssuer{Resolver: s.GetIP}
	s.Register = s
	return s
}

func main() {
	s := InitServer()
	s.PrepServer()
	s.GoServer.Killme = true
	if s.RegisterServer("monitor", false) {
		s.RegisterServingTask(s.emailRunner)
		err := s.Serve()
		if err != nil {
			fmt.Printf("Error serving: %v", err)
		}
	}
}
