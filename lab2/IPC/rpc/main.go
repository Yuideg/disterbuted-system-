package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"github.com/yuidegm/dsp/lab2/IPC/rpc/server"
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		fmt.Println("encountered an error")
		return
	}
	go http.Serve(l, nil)
}
