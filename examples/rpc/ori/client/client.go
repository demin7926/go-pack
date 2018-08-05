package client

import (
	"log"
	"mackee.cn/go-pack/examples/rpc/ori/server"
	"net/rpc"
	"time"
)

func NewClient(proto, address string) (*rpc.Client, error) {
	var client *rpc.Client
	var err error
	if proto == "tcp" {
		client, err = rpc.Dial("tcp", address)
	} else {
		client, err = rpc.DialHTTP("tcp", address)
	}
	if err != nil {
		log.Fatalf("DialHTTP error: %v", err)
		return nil, err
	}
	return client, nil
}

func Calling(client *rpc.Client) error {
	args := server.Args{Num1: 12, Num2: 3}
	var res int

	// call the SimpleRPC.Add by Call()
	log.Println("====================")
	log.Printf(">>> before client.Call: %v", time.Now())
	err := client.Call("SimpleRPC.Add", args, &res)
	log.Printf(">>> after client.Call: %v", time.Now())
	if err != nil {
		log.Fatalf("clent.Call failed: %v", err)
		return err
	}

	// call the SimpleRPC.Add by Go()
	log.Println("====================")
	log.Printf(">>> before client.Go: %v", time.Now())
	asyncCall := client.Go("SimpleRPC.Add", args, &res, nil)
	log.Printf(">>> after client.Go: %v", time.Now())
	replyCall := <-asyncCall.Done
	log.Printf(">>> client.Go complete: %v", time.Now())
	log.Printf("replyCall: %v", replyCall)
	log.Printf("args: %v, result: %v\n", args, res)

	// call the SimpleRPC.Minus by Go()
	log.Println("====================")
	log.Printf(">>> before client.Go: %v", time.Now())
	asyncCall = client.Go("SimpleRPC.Minus", args, &res, nil)
	log.Printf(">>> after client.Go: %v", time.Now())
	replyCall = <-asyncCall.Done
	log.Printf(">>> client.Go complete: %v", time.Now())
	log.Printf("replyCall: %v", replyCall)
	log.Printf("args: %v, result: %v\n", args, res)

	return nil
}
