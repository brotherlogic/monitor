package main

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pbgs "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// ProdIssuer the issuer to use in prod
type ProdIssuer struct {
	Resolver func(string) (string, int)
}

func (p ProdIssuer) createIssue(service, methodCall string, timeMs int32, otherCalls string) {
	ip, port := p.Resolver("githubcard")
	if port > 0 {
		conn, _ := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
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

func (s *Server) emailRunner() {
	for true {
		s.emailSlowFunction()
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
func (s Server) Mote(master bool) error {
	return nil
}

// GetState gets the state of the server
func (s Server) GetState() []*pbgs.State {
	return []*pbgs.State{&pbgs.State{Key: "last_slow", TimeValue: s.LastSlowCheck.Unix()}}
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
