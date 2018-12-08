package main

import (
	"fmt"
	"gofunctional/rpcserver/client"
	"gofunctional/rpcserver/server"
)

func main() {
	go server.StartServer()
	c := client.CreateClient()
	defer c.Close()
	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
