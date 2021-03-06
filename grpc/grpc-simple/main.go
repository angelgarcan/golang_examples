package main

import (
	"context"
	"flag"
	"fmt"

	pb "garcan/examples/grpc-simple/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedParserServer
}

func (s *server) ParseAudit(ctx context.Context, in *pb.ParseMessageRequest) (*pb.ParseMessageResponse, error) {
	log.Printf("Received: %v", in.GetData())
	return &pb.ParseMessageResponse{Response: in.String()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterParserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
