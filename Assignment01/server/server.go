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
	if isValidToken(in.GetToken()) {
		serverTime := time.Now().Unix()
		log.Printf("Received Token: %s [%s]\tSend time: %s", in.GetToken(), "OK", time.Unix(serverTime, 0).String())
		return &pb.TimeReply{
			Success:    true,
			ServerTime: serverTime,
		}, nil
	} else {
		log.Printf("Received Token: %s [%s]", in.GetToken(), "INVALID")
		return &pb.TimeReply{
			Success:    false,
			ServerTime: -1,
		}, nil
	}
}

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	timeServer := grpc.NewServer()
	pb.RegisterTimeServiceServer(timeServer, &server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := timeServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
