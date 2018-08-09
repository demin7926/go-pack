package main

import (
	"github.com/spf13/cobra"
	"log"
	"mackee.cn/go-pack/examples/rpc/ori/client"
	"mackee.cn/go-pack/examples/rpc/ori/server"
	"os"
)

var (
	proto string
	addr  string
)

func runServer(proto, addr string) {
	log.Printf("starting a %s server on %s...", proto, addr)
	if proto == "tcp" {
		server.ServeByTCP(addr)
	} else {
		server.ServeByHTTP(addr)
	}
}

func runClient(proto, addr string) {
	log.Printf("starting a %s client and connect to %s...", proto, addr)
	cli, err := client.NewClient(proto, addr)
	if err != nil {
		log.Fatalf("init client failed: %v", err)
	} else {
		client.Calling(cli)
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "A simple RPC demo",
	}

	serCmd := &cobra.Command{
		Use:   "server [-p PROTO] [-a ADDR]",
		Short: "Run the RPC server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer(proto, addr)
		},
	}
	serCmd.Flags().StringVarP(&proto, "protocol", "p", "TCP", "the transport protocol of rpc, TCP/HTTP")
	serCmd.Flags().StringVarP(&addr, "address", "a", "127.0.0.1:1234", "listenning address")

	cliCmd := &cobra.Command{
		Use:   "client [-p PROTO] [-a ADDR]",
		Short: "Run the RPC client",
		Run: func(cmd *cobra.Command, args []string) {
			runClient(proto, addr)
		},
	}
	cliCmd.Flags().StringVarP(&proto, "protocol", "p", "TCP", "the transport protocol of rpc, TCP/HTTP")
	cliCmd.Flags().StringVarP(&addr, "address", "a", "127.0.0.1:1234", "listenning address")

	rootCmd.AddCommand(serCmd)
	rootCmd.AddCommand(cliCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed in rootCmd.Execute(): %v", err)
		os.Exit(1)
	}
}
