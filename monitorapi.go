package main

import (
	"time"

	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

// WriteMessageLog Writes out a message log
func (s *Server) WriteMessageLog(ctx context.Context, in *pb.MessageLog) (*pb.LogWriteResponse, error) {
	s.writes++
	s.writeMutex.Lock()
	s.writeMap[in.Entry.Name]++
	s.writeMutex.Unlock()
	in.Timestamps = time.Now().Unix()
	s.logs = append(s.logs, in)

	return &pb.LogWriteResponse{Success: true, Timestamp: in.Timestamps}, nil
}

// ReadMessageLogs Reads and returns the message logs for a given entry
func (s *Server) ReadMessageLogs(ctx context.Context, in *pbr.RegistryEntry) (*pb.MessageLogReadResponse, error) {
	s.reads++
	response := &pb.MessageLogReadResponse{Logs: make([]*pb.MessageLog, 0)}
	for _, log := range s.logs {
		if log != nil && log.Entry != nil && (log.Entry.Name == in.GetName() || in.GetName() == "") {
			response.Logs = append(response.Logs, log)
		}
	}
	return response, nil
}
