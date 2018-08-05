package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	Num1, Num2 int
}

type SimpleRPC int

func (t *SimpleRPC) Add(args Args, res *int) error {
	log.Printf("<<< func Add() start: %v", time.Now())
	time.Sleep(time.Second)
	*res = args.Num1 + args.Num2
	log.Printf("<<< func Add() return: %v", time.Now())
	return nil
}

func (t *SimpleRPC) Minus(args Args, res *int) error {
	log.Printf("<<< func Minus() start: %v", time.Now())
	time.Sleep(time.Second * 2)
	*res = args.Num1 - args.Num2
	log.Printf("<<< func Minus() return: %v", time.Now())
	return nil
}

/**
 * 通过HTTP提供服务
 */
func ServeByHTTP(address string) {
	sRPC := new(SimpleRPC)
	rpc.Register(sRPC)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	//更典型的用法是此处用另一个协和启动RPC服务，
	//而主协和继续处理其他事情，但注意不能退出。
	//go http.Serve(l, nil)

	http.Serve(l, nil)
}

/**
 * 通过TCP提供服务
 */
func ServeByTCP(address string) {
	server := rpc.NewServer()
	server.Register(new(SimpleRPC))

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go server.ServeConn(conn)
	}
}
