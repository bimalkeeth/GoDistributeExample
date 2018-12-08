package server

import (
	"fmt"
	"gofunctional/rpcserver/contract"
	"log"
	"net"
	"net/rpc"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	var err = rpc.Register(helloWorld)
	if err != nil {
		log.Fatal("Error in registration")
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %v", port))
	}
	defer func() {
		err := l.Close()
		if err != nil {
			log.Fatal("Error in closing connection")
		}
	}()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error Occurred")
		}
		go rpc.ServeConn(conn)
	}
}

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}
