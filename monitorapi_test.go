package main

import (
	"golang.org/x/net/context"
	"testing"

	pbr "github.com/brotherlogic/discovery/proto"
)

func TestHeartBeat(t *testing.T) {
	s := InitServer()
	r, err := s.ReceiveHeartbeat(context.Background(), &pbr.RegistryEntry{})
	if err != nil {
		t.Errorf("Unable to send heartbeat: %v", err)
	}

	if r.BeatTime <= 0 {
		t.Errorf("Failed to get the right beat time: %v", r.BeatTime)
	}
}
