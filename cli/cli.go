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
		case "logs":
			utils.SendTrace(ctx, "monitorcli-"+os.Args[1]+"-prefind", time.Now(), pbt.Milestone_MARKER, "monitor")
			host, port := findServer("monitor")
			conn, _ := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
			defer conn.Close()
			utils.SendTrace(ctx, "monitorcli-"+os.Args[1]+"-postfind", time.Now(), pbt.Milestone_MARKER, "monitor")

			monitor := pb.NewMonitorServiceClient(conn)
			utils.SendTrace(ctx, "monitorcli-"+os.Args[1]+"-preread", time.Now(), pbt.Milestone_MARKER, "monitor")
			logs, err := monitor.ReadMessageLogs(ctx, &pbdi.RegistryEntry{Name: os.Args[2]})
			utils.SendTrace(ctx, "monitorcli-"+os.Args[1]+"-postread", time.Now(), pbt.Milestone_MARKER, "monitor")
			if err != nil {
				log.Fatalf("Error getting logs: %v", err)
			}
			for _, log := range logs.Logs {
				fmt.Printf("%v. %v\n", time.Unix(log.GetTimestamps(), 0), log)
			}
		}
	}
	utils.SendTrace(ctx, "monitorcli-"+os.Args[1], time.Now(), pbt.Milestone_END, "monitor")
}
