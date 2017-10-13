package main

import (
	"context"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// ProdIssuer the issuer to use in prod
type ProdIssuer struct {
	Resolver func(string) (string, int)
}

func (p ProdIssuer) createIssue(service, methodCall string, timeMs int32) {
	ip, port := p.Resolver("githubcard")
	conn, _ := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	defer conn.Close()
	client := pbgh.NewGithubClient(conn)
	client.AddIssue(context.Background(), &pbgh.Issue{Service: service, Title: "Fix performance", Body: "Fix " + service + " and " + methodCall})
}
func (p ProdIssuer) getSentCount() int {
	return 0
}

const (
	//The number of function call details to keep
	numCalls = 10000
)

func (s Server) emailRunner() {
	for true {
		s.emailSlowFunction()
		time.Sleep(time.Hour * 24)
	}
}

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
	s.write = true
	s.RegisterServer("monitor", true)
	s.RegisterServingTask(s.emailRunner)
	s.Serve()
}
