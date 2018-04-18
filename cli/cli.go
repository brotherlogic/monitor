package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/monitor/monitorproto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

func findServer(name string) (string, int) {
	ip, port, _ := utils.Resolve(name)
	return ip, int(port)
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: build run\n")
	} else {
		switch os.Args[1] {
		case "list":
			host, port := findServer("monitor")

			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()

			monitor := pb.NewMonitorServiceClient(conn)
			stats, err := monitor.GetStats(context.Background(), &pb.FunctionCall{})
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
			logs, err := monitor.ReadMessageLogs(context.Background(), &pbdi.RegistryEntry{})
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
			_, err := monitor.WriteMessageLog(context.Background(), &pb.MessageLog{Message: "Hello"})
			if err != nil {
				log.Fatalf("Error getting logs: %v", err)
			}
		case "clear":
			host, port := findServer("monitor")

			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()

			monitor := pb.NewMonitorServiceClient(conn)
			_, err := monitor.ClearStats(context.Background(), &pb.Empty{})
			if err != nil {
				log.Fatalf("Error getting logs: %v", err)
			}
		}
	}
}
