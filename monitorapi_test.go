package main

import (
	"os"
	"testing"

	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

func InitTestServer() *Server {
	s := InitServer()
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
