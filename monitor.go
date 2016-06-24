package main

func main() {
	s := Server{}
	s.RegisterServer(true)
	s.Serve()
}
