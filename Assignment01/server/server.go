package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	pb "Assignment01/protocol"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 11451, "The server port")
)

type server struct {
	pb.UnimplementedTimeServiceServer
}

func isValidToken(token string) bool {
	return strings.HasPrefix(token, "client_")
}

func (s *server) GetTime(ctx context.Context, in *pb.TimeRequest) (*pb.TimeReply, error) {
	log.Printf("Received Token: %s", in.GetToken())
	if isValidToken(in.GetToken()) {
		return &pb.TimeReply{
			Success:    true,
			ServerTime: time.Now().Unix(),
		}, nil
	} else {
		return &pb.TimeReply{
			Success:    false,
			ServerTime: -1,
		}, nil
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
