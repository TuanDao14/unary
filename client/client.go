package main

import (
	"context"
	"log"

	"github.com/tuanda/unary/unarypb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalf(" err while dial %v", err)
	}
	defer cc.Close()

	client := unarypb.NewCalculatorServiceClient(cc)

	log.Printf("service client %f\n", client)
	callSum(client)

}

func callSum(c unarypb.CalculatorServiceClient) {
	log.Println("calling sum api")
	resp, err := c.Sum(context.Background(), &unarypb.SumRequest{
		Num1: 7,
		Num2: 6,
	})

	if err != nil {
		log.Fatalf("call sum api err %v", err)
	}

	log.Printf("sum api response %v\n", resp.GetResult())
}
