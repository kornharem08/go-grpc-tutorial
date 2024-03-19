package services

import (
	"context"
	"fmt"
)

type CalculatorService interface {
	Hello(name string) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) calculatorService {
	return calculatorService{calculatorClient: calculatorClient}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}

	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	fmt.Printf("Service: %v\n")
	fmt.Printf("Reqest : %v\n", req.Name)
	fmt.Printf("Response: %\n", res.Result)

	return nil
}
