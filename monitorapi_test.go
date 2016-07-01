package main

import (
	"golang.org/x/net/context"
	"testing"
	"time"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

func TestHeatBeatTime(t *testing.T) {
	s := InitServer()
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
	s := InitServer()
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
