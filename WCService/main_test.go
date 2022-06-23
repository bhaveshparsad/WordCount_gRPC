package WCService

import (
	"context"
	pb "WordCount_gRPC/Service"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGetWordCountServer(s, &WCServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestWordCount(t *testing.T) {

	//Dial a connection to grpc Server
	ctx := context.Background()
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer connection.Close()

	//Create new Client
	client := pb.NewGetWordCountClient(connection)

	defer connection.Close()

	//Result
	response, err := client.WordCount(ctx, &pb.Request{Text: "Demo Test for gRPC"})
	if err != nil {
		t.Fatal("Could not count word: \n", err)
	}
	t.Log("WordCount:\n")

	for i,v := range response.Wc{
		t.Logf("Word: %s \t	Count: %d", v.Word, v.Count)
		if i == 9 {
			break
		}
	}
}