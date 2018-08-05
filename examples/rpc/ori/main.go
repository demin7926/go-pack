package main

import (
	"flag"
	"log"

	"mackee.cn/go-pack/examples/rpc/ori/client"
	"mackee.cn/go-pack/examples/rpc/ori/server"
)

func runServer(proto, addr string) {
	if proto == "tcp" {
		server.ServeByTCP(addr)
	} else {
		server.ServeByHTTP(addr)
	}
}

func runClient(proto, addr string) {
	cli, err := client.NewClient(proto, addr)
	if err != nil {
		log.Fatalf("init client failed: %v", err)
	} else {
		client.Calling(cli)
	}
}

func main() {
	mode := flag.String("mode", "server", "running mode, server/client")
	proto := flag.String("proto", "tcp", "transport protocol, tcp/http")
	addr := flag.String("addr", "localhost:1234", "server address")
	flag.Parse()

	log.Printf("running mode: %v, protocol: %v, server address: %v", *mode, *proto, *addr)

	if *mode == "server" {
		runServer(*proto, *addr)
	} else {
		runClient(*proto, *addr)
	}
}
