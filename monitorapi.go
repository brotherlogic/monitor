package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	pbr "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
	pbt "github.com/brotherlogic/tracer/proto"
)

// WriteMessageLog Writes out a message log
func (s *Server) WriteMessageLog(ctx context.Context, in *pb.MessageLog) (*pb.LogWriteResponse, error) {
	if in.Entry == nil || in.Entry.Name == "" {
		s.RaiseIssue(ctx, "Missing Entry", fmt.Sprintf("%v", in.Entry), false)
		return &pb.LogWriteResponse{}, fmt.Errorf("Entry is not specified correctly")
	}
	ctx = s.LogTrace(ctx, "WriteMessageLog", time.Now(), pbt.Milestone_START_FUNCTION)
	s.writes++
	s.writeMutex.Lock()
	s.writeMap[in.Entry.Name]++
	s.writeMutex.Unlock()
	in.Timestamps = time.Now().Unix()
	s.logs = append(s.logs, in)

	//Trim off the logs if we're getting long
	if len(s.logs) > 1000 {
		s.logs = s.logs[500:]
	}

	s.LogTrace(ctx, "WriteMessageLog", time.Now(), pbt.Milestone_END_FUNCTION)
	return &pb.LogWriteResponse{Success: true, Timestamp: in.Timestamps}, nil
}

// ReadMessageLogs Reads and returns the message logs for a given entry
func (s *Server) ReadMessageLogs(ctx context.Context, in *pbr.RegistryEntry) (*pb.MessageLogReadResponse, error) {
	ctx = s.LogTrace(ctx, "ReadMessageLogs", time.Now(), pbt.Milestone_START_FUNCTION)
	s.reads++
	response := &pb.MessageLogReadResponse{Logs: make([]*pb.MessageLog, 0)}
	for _, log := range s.logs {
		if log != nil && log.Entry != nil && (log.Entry.Name == in.GetName() || in.GetName() == "") {
			response.Logs = append(response.Logs, log)
		}
	}

	if len(response.Logs) > 500 {
		response.Logs = response.Logs[len(response.Logs)-500:]
	}

	s.LogTrace(ctx, "ReadMessageLogs", time.Now(), pbt.Milestone_END_FUNCTION)
	return response, nil
}
