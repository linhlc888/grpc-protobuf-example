package main

import (
	"grpc-protobuf-example/pb"
	"fmt"
	"os"
	"log"
	"strconv"
	"google.golang.org/grpc"
	context "golang.org/x/net/context"
)

func main() {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	gcdClient := pb.NewGCDServiceClient(conn)
	a, _ := strconv.ParseUint(os.Args[1], 10, 64)
	b, _ := strconv.ParseUint(os.Args[2], 10, 64)
	req := &pb.GCDRequest{A: a, B: b}
	ctx := context.Background()
	res, err := gcdClient.Compute(ctx, req)
	if err != nil {
		log.Fatalf("Response err : %v", err)
	} else {
		fmt.Printf("GCD of %d and %d is %d.\n", a, b, res.Result)
	}
}
