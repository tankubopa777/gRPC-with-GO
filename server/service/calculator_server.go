package service

import (
	"context"
	"fmt"
)

type calculatorServer struct {

}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (calculatorServer) Hello(ctx context.Context,req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprint("Hello ", req.Name)
	rest := HelloResponse{
		Result: result,
	}
	return &rest, nil
}