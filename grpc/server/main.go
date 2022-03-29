/*
@Time : 2022/2/11 15:38
@Author : yuanhai
*/

package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	pb "voyage/grpc/helloworld"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "Hello " + in.GetGreeting()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
