package main

import (
	"math/rand"
	"testing"

	"golang.org/x/net/context"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/proto"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestMemoryUsage(t *testing.T) {
	s := InitTestServer()

	for i := 0; i < 100000; i++ {
		strfuzz := randStringBytes(10)
		s.WriteMessageLog(context.Background(), &pb.MessageLog{Entry: &pbd.RegistryEntry{Identifier: "Blah" + strfuzz, Name: "Test" + strfuzz}})
	}
}
