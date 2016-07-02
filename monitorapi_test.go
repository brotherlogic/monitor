package main

import (
	"golang.org/x/net/context"
	"os"
	"strconv"
	"testing"
	"time"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

func InitTestServer() Server {
	s := InitServer()
	s.logDirectory = "testlogs"
	return s
}

func TestWriteValueLog(t *testing.T) {
	s := InitTestServer()

	registry := &pbr.RegistryEntry{Identifier: "Blah", Name: "Test"}
	valueLog := &pb.ValueLog{Entry: registry, Value: 35.5}
	r, err := s.WriteValueLog(context.Background(), valueLog)
	if err != nil {
		t.Errorf("Write Value Log has returned an error")
	}

	//Check that the log has been written in the required directory
	if _, err := os.Stat("testlogs/Test/Blah/" + strconv.Itoa(int(r.Timestamp)) + ".value"); os.IsNotExist(err) {
		t.Errorf("Logs do not exist")
	}
}

func TestWriteMessageLog(t *testing.T) {
	s := InitTestServer()

	registry := &pbr.RegistryEntry{Identifier: "Blah", Name: "Test"}
	messageLog := &pb.MessageLog{Entry: registry, Message: "This is the log message"}
	r, err := s.WriteMessageLog(context.Background(), messageLog)
	if err != nil {
		t.Errorf("Write Value Log has returned an error")
	}

	//Check that the log has been written in the required directory
	if _, err := os.Stat("testlogs/Test/Blah/" + strconv.Itoa(int(r.Timestamp)) + ".message"); os.IsNotExist(err) {
		t.Errorf("Logs do not exist")
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
