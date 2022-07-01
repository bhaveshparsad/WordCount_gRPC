package main

import (
	"bufio"
	"context"
	"fmt"
	pb "WordCount_gRPC/protoFiles"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	
	
)

const address = "localhost:60020"

func main() {
	fmt.Println("Input String :")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}


	// Connecting to grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal("Did not connect", err)
	}

	defer conn.Close()

	//Creating new client
	client := pb.NewGetWordCountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Result
	response, err := client.WordCount(ctx, &pb.Request{Text: str})
	if err != nil {
		log.Fatal("Error: \n", err)
	}


	for i, v := range response.Wc {
		log.Printf("Word: %s \t Count: %d", v.Word, v.Count)
		if i == 9 {
			break
		}
	}
}
