package main

import (
	"context"
	"fmt"
	"log"
	"sabio-ekuator/config"
	"sabio-ekuator/pb"
)

func (s *Server) CreateProduct(ctx context.Context, req *pb.ProductReq) (*pb.ProductMsg, error) {
	fmt.Printf("\nCreate product was invoked with %v", req)

	query := `
	INSERT INTO Product (name, price, stock)
	VALUES ($1, $2, $3)
	`
	_, err := config.DB.Query(query, req.Name, req.Price, req.Stock)
	if err != nil {
		log.Fatalf("\nErr querying create product: %v", err)
	}

	return &pb.ProductMsg{
		Message: "New product created",
	}, nil
}
