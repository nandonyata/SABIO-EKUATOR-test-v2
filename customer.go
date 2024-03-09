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
	query := `
	SELECT id, name, email FROM Customer
	ORDER BY id DESC
	`

	row, err := config.DB.Query(query)
	if err != nil {
		status.Error(codes.Internal, fmt.Sprintf("Err: %v\n", err))
	}

	for row.Next() {
		cust := entity.Customer{}

		if err = row.Scan(&cust.Id, &cust.Name, &cust.Email); err != nil {
			status.Error(codes.Internal, fmt.Sprintf("Err scanning: %v\n", err))
		}

		stream.Send(entity.DocToCustomer(&cust))
	}

	return nil
}

func (s *Server) FethcOneCustomer(_ context.Context, req *pb.CustomerId) (*pb.Customer, error) {
	query := `
	SELECT id, name, email FROM Customer
	WHERE id = $1
	`

	cust := entity.Customer{}
	err := config.DB.QueryRow(query, req.Id).Scan(&cust.Id, &cust.Name, &cust.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Err : %v\n", err))
	}

	return entity.DocToCustomer(&cust), nil
}

func (s *Server) UpdateCustomer(_ context.Context, req *pb.Customer) (*pb.CustomerMsg, error) {
	query := `
	UPDATE Customer
	SET name = $2, email = $3
	WHERE id = $1
	`
	_, err := config.DB.Exec(query, req.Id, req.Name, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Err: %v\n", err))
	}

	return &pb.CustomerMsg{
		Message: "Success update customer",
	}, nil
}

func (s *Server) DeleteCustomer(_ context.Context, req *pb.CustomerId) (*pb.CustomerMsg, error) {
	query := `
		DELETE FROM Customer
		WHERE id = $1
	`

	config.DB.Exec(query, req.Id)

	return &pb.CustomerMsg{
		Message: "Succes delete customer",
	}, nil
}
