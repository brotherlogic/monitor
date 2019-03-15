package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

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
	ctx, cancel := utils.BuildContext(fmt.Sprintf("monitorcli-%v", os.Args[1]), "monitor")
	defer cancel()

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: build run\n")
	} else {
		switch os.Args[1] {
		case "logs":
			host, port := findServer("monitor")
			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()

			monitor := pb.NewMonitorServiceClient(conn)
			logs, err := monitor.ReadMessageLogs(ctx, &pbdi.RegistryEntry{Name: os.Args[2]})
			if err != nil {
				log.Fatalf("Error getting logs: %v", err)
			}

			sort.SliceStable(logs.Logs, func(i, j int) bool {
				return logs.Logs[i].GetTimestamps() > logs.Logs[j].GetTimestamps()
			})

			for _, log := range logs.Logs {
				fmt.Printf("%v (%v). %v\n", time.Unix(log.GetTimestamps(), 0), log.Entry.Identifier, log.Message)
			}
		}
	}
}
