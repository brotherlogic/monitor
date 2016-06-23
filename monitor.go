package main

import (
	"github.com/brotherlogic/goserver"
)

// MonitorServer the main server type
type MonitorServer struct {
	goserver.GoServer
}

func main() {
	s := MonitorServer{}
	s.RegisterServer(true)
	s.Serve()
}
