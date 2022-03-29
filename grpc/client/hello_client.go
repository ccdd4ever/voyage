/*
@Time : 2022/2/11 17:21
@Author : yuanhai
*/

package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "voyage/grpc/helloworld"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connect to %s failed: %v", address, err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	name := defaultName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Greeting: name})
	if err != nil {
		log.Fatalf("call SayHello failed:%v", err)
	}
	log.Printf("call SayHello resp:%s", r.GetReply())

}
