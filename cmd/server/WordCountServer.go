package main

import (
	"log"
	"net"
	pb "WordCount_gRPC/protoFiles"
	"WordCount_gRPC/wordCountMain"
	"google.golang.org/grpc"
)

const (
	port = ":60020"
)

func main () {
	lis,err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to Listen", err)
	}

	gRPCServer := grpc.NewServer() // Initializing new server

	pb.RegisterGetWordCountServer(gRPCServer, &wordCountMain.WCServer{}) // Registering as new grpc service

	log.Println("Server Listening", lis.Addr())

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatal("Failed to Serve", err)
	}
}

