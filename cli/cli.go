package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/monitor/monitorproto"
)

func findServer(name string) (string, int) {
	conn, err := grpc.Dial("192.168.86.64:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot reach discover server: %v (trying to discover %v)", err, name)
	}
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	rs, err := registry.ListAllServices(context.Background(), &pbdi.Empty{})

	if err != nil {
		log.Fatalf("Failure to list: %v", err)
	}

	for _, r := range rs.Services {
		if r.Name == name {
			log.Printf("%v -> %v", name, r)
			return r.Ip, int(r.Port)
		}
	}

	log.Printf("No Cardserver running")

	return "", -1
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
			stats, err := monitor.GetStats(context.Background(), &pb.FunctionCall{Binary: "discogssyncer", Name: "GetSingleRelease"})
			if err != nil {
				log.Printf("Error getting stats: %v", err)
			}
			fmt.Printf("Stats: %v\n", stats)
		}
	}
}
