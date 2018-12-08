package server

import (
	"GoDistributeExample/httprpcserver/contract"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
func StartServer() {
	helloworld := &HelloWorldHandler{}
	err := rpc.Register(helloworld)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in registering with %v", port))
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in listning with %v", port))
	}
	log.Printf("Server starting %v", port)
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in listning with %v", port))
	}
}
