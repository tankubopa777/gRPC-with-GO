package main

import (
	"client/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()


	calculatorClient := service.NewCalculatorClient(cc)
	calculatorService := service.NewCalculatorService(calculatorClient)

	err = calculatorService.Hello("Tansan")
	if err != nil {
		log.Fatalf("Error while calling Hello RPC: %v", err)
	}
}
