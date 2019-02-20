package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"study/rpcdemo/server"
)

func main() {
	conn,err := net.Dial("tcp",":1213")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := jsonrpc.NewClient(conn)
	defer client.Close()
	var res float64
	err = client.Call("DemoService.Div",server.Args{3,4},&res)
	if err != nil {
		panic(err)
	}
	fmt.Println("res is ",res)

}
