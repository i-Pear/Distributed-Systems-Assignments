package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "Assignment01/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultAddress = "localhost:11451"
	defaultToken   = "client"
)

var (
	addr  = flag.String("addr", defaultAddress, "the address to connect to")
	token = flag.String("token", defaultToken, "token for authorization")
)

func getToken() string {
	return "client_" + strconv.Itoa(rand.Int())
}

func main() {
	flag.Parse()
	// set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTimeServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for true {
		response, err := client.GetTime(ctx, &pb.TimeRequest{
			Token: getToken(),
		})
		if err != nil {
			fmt.Printf("could not get time: %v\n", err)
		} else {
			timestr := time.Unix(response.GetServerTime(), 0).String()
			fmt.Printf("Time from server: %s\n", timestr)
		}
		time.Sleep(time.Second)
	}
}
