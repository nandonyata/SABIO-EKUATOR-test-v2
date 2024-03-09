package main

import (
	"context"
	"fmt"
	"sabio-ekuator/config"
	"sabio-ekuator/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCustomer(_ context.Context, req *pb.Customer) (*pb.CustomerMsg, error) {
	query := `
		INSERT INTO Customer (name, email)
		VALUES ($1, $2)
	`

	if req.Email == "" || req.Name == "" {
		return nil, status.Error(codes.Canceled, "Fill All Field")
	}

	_, err := config.DB.Exec(query, req.Name, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Err: %v\n", err))
	}

	return &pb.CustomerMsg{
		Message: "Success create customer",
	}, nil
}

func (s *Server) FetchAllCustomer(_ *pb.EmptyCustomer, stream pb.CustomerService_FetchAllCustomerServer) error {
	return nil
}

func (s *Server) FethcOneCustomer(context.Context, *pb.CustomerId) (*pb.Customer, error) {
	return &pb.Customer{}, nil
}

func (s *Server) UpdateCustomer(context.Context, *pb.Customer) (*pb.CustomerMsg, error) {
	return &pb.CustomerMsg{}, nil
}

func (s *Server) DeleteCustomer(context.Context, *pb.CustomerId) (*pb.CustomerMsg, error) {
	return &pb.CustomerMsg{}, nil
}
