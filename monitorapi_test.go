package main

import (
	"os"
	"testing"

	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/proto"
)

func InitTestServer() *Server {
	s := InitServer()
	s.logDirectory = "testlogs"
	os.RemoveAll(s.logDirectory)
	s.GoServer.KSclient = *keystoreclient.GetTestClient("./testing")
	s.SkipLog = true
	s.SkipIssue = true
	return s
}

func TestReadMessageLogs(t *testing.T) {
	s := InitTestServer()

	registry := &pbr.RegistryEntry{Identifier: "Blah", Name: "Test"}
	messageLog := &pb.MessageLog{Entry: registry, Message: "This is the log message", Level: pb.LogLevel_KEEP_WEEK}
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

func TestFailWrite(t *testing.T) {
	s := InitTestServer()

	messageLog := &pb.MessageLog{Message: "This is the log message", Level: pb.LogLevel_KEEP_WEEK}
	_, err := s.WriteMessageLog(context.Background(), messageLog)
	if err == nil {
		t.Errorf("Write Value Log has not returned an error")
	}
}

func TestReadMessageLogsTrunc(t *testing.T) {
	s := InitTestServer()

	registry := &pbr.RegistryEntry{Identifier: "Blah", Name: "Test"}
	messageLog := &pb.MessageLog{Entry: registry, Message: "This is the log message", Level: pb.LogLevel_KEEP_WEEK}
	for i := 0; i < 1000; i++ {
		_, err := s.WriteMessageLog(context.Background(), messageLog)
		if err != nil {
			t.Errorf("Write Value Log has returned an error")
		}
	}

	logs, err := s.ReadMessageLogs(context.Background(), registry)
	if err != nil {
		t.Errorf("Error returned on read message logs %v", err)
	}
	if len(logs.Logs) == 1000 {
		t.Errorf("Too many logs: %v", len(logs.Logs))
	}

	if logs.Logs[0].Message != "This is the log message" {
		t.Errorf("Read log the wrong: %v", logs.Logs[0])
	}
}
