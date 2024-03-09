package main

import (
	"context"
	"fmt"
	"log"
	"sabio-ekuator/config"
	"sabio-ekuator/entity"
	"sabio-ekuator/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) GetAllProduct(_ *pb.ProductEmpty, stream pb.ProductService_GetAllProductServer) error {
	fmt.Println("Get all product was invoked")

	query := `
	SELECT name, price, stock FROM Product
	ORDER BY id DESC
	`

	row, err := config.DB.Query(query)
	if err != nil {
		return status.Error(codes.Internal, "internal server error")
	}

	for row.Next() {
		product := &entity.Product{}

		if err := row.Scan(&product.Name, &product.Price, &product.Stock); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("error scanning: %v\n", err))
		}

		stream.Send(entity.DocToProduct(product))
	}

	return nil
}
