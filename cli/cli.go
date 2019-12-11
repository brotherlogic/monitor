package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/monitor/proto"

	//Needed to pull in gzip encoding init

	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&utils.DiscoveryClientResolverBuilder{})
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: build run\n")
	} else {
		switch os.Args[1] {
		case "logs":
			conn, err := grpc.Dial("discovery:///monitor", grpc.WithInsecure(), grpc.WithBalancerName("my_pick_first"))
			if err != nil {
				log.Fatalf("Dial error: %v", err)
			}
			defer conn.Close()

			var logs *pb.MessageLogReadResponse
			err = fmt.Errorf("An error")
			count := 0
			for err != nil && count < 4 {
				ctx, cancel := utils.ManualContext(fmt.Sprintf("monitorcli-%v", os.Args[1]), "monitor", time.Minute)
				defer cancel()

				monitor := pb.NewMonitorServiceClient(conn)
				logs, err = monitor.ReadMessageLogs(ctx, &pbdi.RegistryEntry{Name: os.Args[2]})
				if err != nil {
					fmt.Printf("%v, Error getting logs: %v\n", count, err)
				}
				fmt.Printf("Count: %v -> %v\n", count, err)
				count++
			}
			fmt.Printf("Source -> %v [%v]\n", logs.Server, len(logs.Logs))

			sort.SliceStable(logs.Logs, func(i, j int) bool {
				return logs.Logs[i].GetTimestamps() > logs.Logs[j].GetTimestamps()
			})

			for _, log := range logs.Logs {
				fmt.Printf("%v (%v). %v\n", time.Unix(log.GetTimestamps(), 0), log.Entry.Identifier, log.Message)
			}
		}
	}
}
