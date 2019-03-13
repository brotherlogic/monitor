package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

type logHolder struct {
	logs    []*pb.MessageLog
	pointer int
}

// WriteMessageLog Writes out a message log
func (s *Server) WriteMessageLog(ctx context.Context, in *pb.MessageLog) (*pb.LogWriteResponse, error) {
	if in.Entry == nil || in.Entry.Name == "" {
		s.RaiseIssue(ctx, "Missing Entry", fmt.Sprintf("%v", in.Entry), false)
		return &pb.LogWriteResponse{}, fmt.Errorf("Entry is not specified correctly")
	}
	s.writes++
	s.writeMutex.Lock()
	s.writeMap[in.Entry.Name]++
	s.writeMutex.Unlock()
	in.Timestamps = time.Now().Unix()

	s.logsMutex.Lock()
	if _, ok := s.logs[in.Entry.Name]; !ok {
		s.logs[in.Entry.Name] = &logHolder{
			logs:    make([]*pb.MessageLog, 200),
			pointer: 0,
		}
	}
	s.logs[in.Entry.Name].logs[s.logs[in.Entry.Name].pointer] = in
	s.logs[in.Entry.Name].pointer++
	s.logs[in.Entry.Name].pointer %= 200
	s.logsMutex.Unlock()

	return &pb.LogWriteResponse{Success: true, Timestamp: in.Timestamps}, nil
}

// ReadMessageLogs Reads and returns the message logs for a given entry
func (s *Server) ReadMessageLogs(ctx context.Context, in *pbr.RegistryEntry) (*pb.MessageLogReadResponse, error) {
	s.reads++
	response := &pb.MessageLogReadResponse{Logs: make([]*pb.MessageLog, 0)}
	if val, ok := s.logs[in.Name]; ok {
		for _, l := range val.logs {
			if l != nil {
				response.Logs = append(response.Logs, l)
			}
		}
	}

	if len(response.Logs) > 500 {
		response.Logs = response.Logs[len(response.Logs)-500:]
	}

	return response, nil
}
