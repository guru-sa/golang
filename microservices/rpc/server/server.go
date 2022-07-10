package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"rpc/contarct"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contarct.HelloWorldRequest, reply *contarct.HelloWorldResponse) error {
	reply.Message = fmt.Sprintf("Hello %s", args.Name)
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
