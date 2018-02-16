package main

import (
	"os"
	"testing"
	"time"

	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

type testIssuer struct {
	count int
}

func (t *testIssuer) createIssue(service, methodCall string, timeMs int32, other string) {
	t.count++
}
func (t *testIssuer) getSentCount() int {
	return t.count
}

func InitTestServer() *Server {
	s := InitServer()
	s.issuer = &testIssuer{count: 0}
	s.logDirectory = "testlogs"
	os.RemoveAll(s.logDirectory)
	return s
}

func TestReadMessageLogs(t *testing.T) {
	s := InitTestServer()

	registry := &pbr.RegistryEntry{Identifier: "Blah", Name: "Test"}
	messageLog := &pb.MessageLog{Entry: registry, Message: "This is the log message"}
	_, err := s.WriteMessageLog(context.Background(), messageLog)
	if err != nil {
		t.Errorf("Write Value Log has returned an error")
	}

	logs, err := s.ReadMessageLogs(context.Background(), registry)
	if err != nil {
		t.Errorf("Error returned on read message logs %v", err)
	}
	if len(logs.Logs) != 1 {
		t.Errorf("Too many logs: %v", len(logs.Logs))
	}

	if logs.Logs[0].Message != "This is the log message" {
		t.Errorf("Read log the wrong: %v", logs.Logs[0])
	}
}

func TestClearStats(t *testing.T) {
	s := InitTestServer()
	s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 10})
	stats, _ := s.GetStats(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction"})
	if len(stats.GetStats()) == 0 {
		t.Fatalf("Stats not added!: %v", t)
	}
	s.ClearStats(context.Background(), &pb.Empty{})
	stats, _ = s.GetStats(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction"})
	if len(stats.GetStats()) != 0 {
		t.Fatalf("Stats not cleared!: %v", t)
	}
}

func TestComputingMeanOnFunctionCalls(t *testing.T) {
	s := InitTestServer()

	_, err := s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 10})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}
	_, err = s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 20})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}
	_, err = s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 30})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}

	stats, err := s.GetStats(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction"})
	if err != nil {
		t.Errorf("Failure to produce stats: %v", err)
	}

	if stats.Stats[0].NumberOfCalls != 3 || stats.Stats[0].MeanRunTime != 20 {
		t.Errorf("Stats have come back wrong: %v", stats)
	}
}

func TestWriteEmptyFunctionCall(t *testing.T) {
	s := InitTestServer()

	_, err := s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 0})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}
}

func TestMonitorFunctionCalls(t *testing.T) {
	s := InitTestServer()

	_, err := s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 340})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}
	_, err = s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 340})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}

	stats, err := s.GetStats(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction"})
	if err != nil {
		t.Errorf("Failure to produce stats: %v", err)
	}

	if stats.Stats[0].NumberOfCalls != 2 || stats.Stats[0].MeanRunTime != 340 {
		t.Errorf("Stats have come back wrong: %v", stats)
	}
}

func TestEmailFunctionCalls(t *testing.T) {
	s := InitTestServer()
	_, err := s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RPCRunFunction", Time: 900})
	if err != nil {
		t.Fatalf("Failed to write the function call %v", err)
	}

	s.emailSlowFunction()
	if s.issuer.getSentCount() != 1 {
		t.Errorf("Failed to get the count of sent")
	}
}

func TestMonitorAll(t *testing.T) {
	s := InitTestServer()

	_, err := s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 340})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}
	_, err = s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 340})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}

	stats, err := s.GetStats(context.Background(), &pb.FunctionCall{})
	if err != nil {
		t.Errorf("Failure to produce stats: %v", err)
	}

	if stats.Stats[0].NumberOfCalls != 2 || stats.Stats[0].MeanRunTime != 340 || len(stats.Stats[0].RunTimes) != 2 {
		t.Errorf("Stats have come back wrong: %v", stats)
	}
}

func TestTrimFunctionCalls(t *testing.T) {
	s := InitTestServer()

	for i := 0; i < 20000; i++ {
		_, err := s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 340})
		if err != nil {
			t.Errorf("Failure to write the function call: %v", err)
		}
	}

	stats, err := s.GetStats(context.Background(), &pb.FunctionCall{})
	if err != nil {
		t.Errorf("Failure to produce stats: %v", err)
	}

	if stats.Stats[0].NumberOfCalls != 20000 || stats.Stats[0].MeanRunTime != 340 || len(stats.Stats[0].RunTimes) != 10000 {
		t.Errorf("Stats have come back wrong: %v, %v, %v", stats.Stats[0].NumberOfCalls, stats.Stats[0].MeanRunTime, len(stats.Stats[0].RunTimes))
	}
}

func TestPullNonFunction(t *testing.T) {
	s := InitTestServer()

	_, err := s.WriteFunctionCall(context.Background(), &pb.FunctionCall{Binary: "madeup", Name: "RunFunction", Time: 340})
	if err != nil {
		t.Errorf("Failure to write the function call: %v", err)
	}

	stats, err := s.GetStats(context.Background(), &pb.FunctionCall{Binary: "madeupbinary", Name: "RunFunction"})
	if err == nil {
		t.Errorf("Stats have come back even though they shouldn't: %v", stats)
	}
}

func TestHeatBeatTime(t *testing.T) {
	s := InitTestServer()
	r, _ := s.ReceiveHeartbeat(context.Background(), &pbr.RegistryEntry{Identifier: "Blah", Name: "Test"})

	firstTime := r.BeatTime
	time.Sleep(1 * time.Second)
	s.ReceiveHeartbeat(context.Background(), &pbr.RegistryEntry{Identifier: "Blah", Name: "Test"})

	r3, _ := s.GetHeartbeats(context.Background(), &pb.Empty{})
	if len(r3.Beats) != 1 {
		t.Errorf("Too many heartbeats")
	}

	if r3.GetBeats()[0].BeatTime == firstTime {
		t.Errorf("Time has not updated")
	}

	if r3.GetBeats()[0].Entry.Name != "Test" {
		t.Errorf("Name is not being returned %v", r3.GetBeats()[0])
	}
}

func TestHeartBeat(t *testing.T) {
	s := InitTestServer()
	r, err := s.ReceiveHeartbeat(context.Background(), &pbr.RegistryEntry{})
	if err != nil {
		t.Errorf("Unable to send heartbeat: %v", err)
	}

	if r.BeatTime <= 0 {
		t.Errorf("Failed to get the right beat time: %v", r.BeatTime)
	}

	r2, err := s.GetHeartbeats(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("Unable to get list of heartbeats: %v", err)
	}

	if len(r2.Beats) == 0 {
		t.Errorf("No heartbeats received")
	}
}
