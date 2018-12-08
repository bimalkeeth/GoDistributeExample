package main

import (
	"GoDistributeExample/httprpcserver/client"
	"GoDistributeExample/httprpcserver/server"
	"fmt"
)

func main() {
	server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
