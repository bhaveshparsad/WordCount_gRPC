package WCService

import (
	"context"
	"log"
	"sort"
	"strings"
	pb "WordCount_gRPC/Service"
)

type WCServer struct {
	pb.UnimplementedGetWordCountServer
}

func (s *WCServer) WordCount(ctx context.Context, in *pb.Request) (*pb.Response, error){

	log.Println("Received")
	str := in.Text

	count := make(map[string]int)

	for _, word := range strings.Fields(str) {
		count[word]++ // Counting words
	}

	// words := make([]string, 0, len(count))
	// for i := range count {
	// 	words = append(words, i) // Appending words and their count
	// }

	var WC []*pb.WordCount// Slice of WordCount type
	for k, v := range count{
		WC = append(WC, &pb.WordCount{Word: k, Count: uint32(v)}) // Appending words and their count in slice
	}

	sort.Slice(WC, func(i, j int) bool {
		return WC[i].Count > WC[j].Count // Sorting 
	})

	return &pb.Response{Wc : WC}, nil

}