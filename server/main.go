package main

import (
	"fmt"
	"log"
	"net"
	"server/service"

	"google.golang.org/grpc"
)

func main(){
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Unable to listen on port 50051: %v", err)
	}

	// Register the service with the server
	service.RegisterCalculatorServer(s, service.NewCalculatorServer())

	
	fmt.Println("Server is running on port 50051")
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Unable to serve: %v", err)
	}
}