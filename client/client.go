package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/radicalbit/NSDb-go/proto"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/testdata"
)

var (
	serverAddr = flag.String("address", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {

	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHealthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Check(ctx, &pb.HealthCheckRequest{Service: "whatever"})
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}

	log.Printf("STATUS: %s", r.GetStatus())
}
