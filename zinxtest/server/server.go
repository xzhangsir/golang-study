package main

import "zinx/znet"

func main() {
	s := znet.NewServer("zx")
	s.Serve()
}
