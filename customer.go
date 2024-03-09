package main

import (
	"context"
	"sabio-ekuator/pb"
)

func (s *Server) CreateCustomer(_ context.Context, req *pb.Customer) (*pb.CustomerMsg, error) {

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
