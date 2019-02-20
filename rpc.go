package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"study/rpcdemo/server"
)

//telnet localhost 1213
//{"method":"DemoService.Div","Params":[{"A":8,"B":6}],"id":12}
//{"id":12,"result":1.3333333333333333,"error":null}
func main() {
	rpc.Register(server.DemoService{})
	listener,err := net.Listen("tcp",":1213")
	if err != nil {
		panic(err)
	}
	for  {
		conn,err := listener.Accept()
		if err != nil {
			log.Printf("%v\n",err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
