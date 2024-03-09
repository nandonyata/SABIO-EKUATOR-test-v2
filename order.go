package main

import (
	"context"
	"sabio-ekuator/pb"
)

func (s *Server) CreateOrder(_ context.Context, req *pb.Order) (*pb.OrderMessage, error) {
	return &pb.OrderMessage{}, nil
}

func (s *Server) FetchOneOrder(_ context.Context, req *pb.OrderId) (*pb.Order, error) {
	return &pb.Order{}, nil
}

func (s *Server) FetchAllOrder(_ *pb.OrderEmpty, stream pb.OrderService_FetchAllOrderServer) error {
	return nil
}
