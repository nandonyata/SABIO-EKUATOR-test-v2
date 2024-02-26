package main

import (
	"context"
	"fmt"
	"log"
	"sabio-ekuator/config"
	"sabio-ekuator/entity"
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

func (s *Server) FetchOneProduct(ctx context.Context, req *pb.ProductId) (*pb.ProductReq, error) {
	fmt.Println("Fetch one product was invoked with ", req)

	var res entity.Product

	query := `
	SELECT name, price, stock from Product
	WHERE id = $1
	`

	config.DB.QueryRow(query, req.Id).Scan(&res.Name, &res.Price, &res.Stock)

	if res.Name == "" && res.Price == 0 && res.Stock == 0 {
		log.Fatalf("Data not found with id %v", req.Id)
	}

	return &pb.ProductReq{
		Name:  res.Name,
		Price: res.Price,
		Stock: res.Stock,
	}, nil
}
