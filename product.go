package main

import (
	"context"
	"fmt"
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
		return nil, status.Error(codes.Internal, fmt.Sprintf("Err querying create product: %v", err))
	}

	return &pb.ProductMsg{
		Message: "New product created",
	}, nil
}

func (s *Server) FetchOneProduct(ctx context.Context, req *pb.ProductId) (*pb.ProductReq, error) {
	fmt.Println("Fetch one product was invoked with ", req)

	var res entity.Product

	query := `
	SELECT name, price, stock, id from Product
	WHERE id = $1
	`

	config.DB.QueryRow(query, req.Id).Scan(&res.Name, &res.Price, &res.Stock, &res.Id)

	if res.Name == "" && res.Price == 0 && res.Stock == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Data not found with id %v", req.Id))
	}

	return &pb.ProductReq{
		Id:    res.Id,
		Name:  res.Name,
		Price: res.Price,
		Stock: res.Stock,
	}, nil
}

func (s *Server) FetchAllProduct(_ *pb.ProductEmpty, stream pb.ProductService_FetchAllProductServer) error {
	fmt.Println("Get all product was invoked")

	query := `
	SELECT name, price, stock, id FROM Product
	ORDER BY id DESC
	`

	row, err := config.DB.Query(query)
	if err != nil {
		return status.Error(codes.Internal, "internal server error")
	}

	for row.Next() {
		product := &entity.Product{}

		if err := row.Scan(&product.Name, &product.Price, &product.Stock, &product.Id); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("error scanning: %v\n", err))
		}

		stream.Send(entity.DocToProduct(product))
	}

	return nil
}

func (s *Server) UpdateProduct(ctx context.Context, req *pb.ProductReq) (*pb.ProductMsg, error) {
	fmt.Println("Update product was invoked with ", req)

	query := `
	UPDATE Product
	SET name = $2, price = $3, stock = $4
	WHERE id = $1
	`

	if _, err := config.DB.Exec(query, req.Id, req.Name, req.Price, req.Stock); err != nil {
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	return &pb.ProductMsg{
		Message: "Success update",
	}, nil
}

func (s *Server) DeleteProduct(_ context.Context, req *pb.ProductId) (*pb.ProductMsg, error) {
	fmt.Println("Delete product was invoked with ", req)

	_, err := config.DB.Exec(`
		DELETE FROM Product
		WHERE id = $1
	`, req.Id)

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error: %v\n", err))
	}

	return &pb.ProductMsg{
		Message: "Success delete product",
	}, nil
}
