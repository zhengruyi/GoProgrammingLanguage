package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp","localhost:1234")
	if err != nil {
		log.Fatal("dialing:",err)
	}
	var reply string
	//通过client.Call来进行rpc的远程调用
	err = client.Call("HelloService.Hello"," ruyi",&reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
