package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	pbgs "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/monitor/monitorproto"
	pbt "github.com/brotherlogic/tracer/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

func findServer(name string) (string, int) {
	ip, port, _ := utils.Resolve(name)
	return ip, int(port)
}

func main() {
	ctx, cancel := utils.BuildContext("monitorcli-"+os.Args[1], "monitor", pbgs.ContextType_MEDIUM)
	defer cancel()

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: build run\n")
	} else {
		switch os.Args[1] {
		case "list":
			host, port := findServer("monitor")

			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()

			monitor := pb.NewMonitorServiceClient(conn)
			stats, err := monitor.GetStats(ctx, &pb.FunctionCall{})
			if err != nil {
				log.Fatalf("Error getting stats: %v", err)
			}
			for _, stat := range stats.Stats {
				fmt.Printf("Stats: %v -> %v (%v)\n", stat.Binary+"-"+stat.Name, stat.MeanRunTime, stat.NumberOfCalls)
				for _, milestone := range stat.Slowest.Milestones {
					fmt.Printf("  Milestone %v -> %v\n", milestone.Name, milestone.Time)
				}
			}
		case "logs":
			host, port := findServer("monitor")

			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()

			monitor := pb.NewMonitorServiceClient(conn)
			logs, err := monitor.ReadMessageLogs(ctx, &pbdi.RegistryEntry{Name: os.Args[2]})
			if err != nil {
				log.Fatalf("Error getting logs: %v", err)
			}
			for _, log := range logs.Logs {
				fmt.Printf("%v. %v\n", time.Unix(log.GetTimestamps(), 0), log)
			}
		case "write":
			host, port := findServer("monitor")

			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()

			monitor := pb.NewMonitorServiceClient(conn)
			_, err := monitor.WriteMessageLog(ctx, &pb.MessageLog{Message: "Hello"})
			if err != nil {
				log.Fatalf("Error getting logs: %v", err)
			}
		case "clear":
			host, port := findServer("monitor")

			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()

			monitor := pb.NewMonitorServiceClient(conn)
			_, err := monitor.ClearStats(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Error getting logs: %v", err)
			}
		}
	}
	utils.SendTrace(ctx, "monitorcli-"+os.Args[1], time.Now(), pbt.Milestone_END, "monitor")
}
