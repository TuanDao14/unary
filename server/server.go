package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tuanda/unary/unarypb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *unarypb.SumRequest) (*unarypb.SumResponse, error) {
	log.Println("sum called...")
	resp := &unarypb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	unarypb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("calculator is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
