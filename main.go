package main

import (
	"flag"
	"fmt"
	"runtime"

	"./server"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8001"
	CONN_TYPE = "tcp"
)

func main() {

	ROOT := flag.String("r", "/home/ilia/go/src/github.com/iliabulavintsev/server/", " dir Root")
	NUM_CPU := flag.Int("c", 1, "num of cpu")

	flag.Parse()

	fmt.Println("root: ", *ROOT)
	fmt.Println("cpu: ", *NUM_CPU)
	runtime.GOMAXPROCS(*NUM_CPU)

	server := server.Server{}
	server.Create(CONN_HOST, CONN_PORT, *ROOT)
	server.Start()

}
