package main

import (
	"context"
	"fmt"
	"sabio-ekuator/pb"
)

func (s *Server) CreateProduct(ctx context.Context, req *pb.ProductReq) (*pb.ProductMsg, error) {
	fmt.Printf("Create product was invoked with %v", req)

	return &pb.ProductMsg{
		Message: "WOII",
	}, nil
}
