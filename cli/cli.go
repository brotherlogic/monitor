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
	conn, err := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot reach discover server: %v (trying to discover %v)", err, name)
	}
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := registry.Discover(ctx, &pbdi.RegistryEntry{Name: "monitor"})

	if err != nil {
		log.Fatalf("Failure to list: %v", err)
	}

	return r.Ip, int(r.Port)
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
				fmt.Printf("%v\n", log)
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
		}
	}
}
