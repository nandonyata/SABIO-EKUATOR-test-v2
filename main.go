package main

import (
	"fmt"
	"log"
	"net"
	"sabio-ekuator/config"
	"sabio-ekuator/pb"

	"google.golang.org/grpc"
)

var PORT = "3000"

type Server struct {
	pb.ProductServiceServer
	pb.CustomerServiceServer
}

func main() {
	config.ConnectDB()

	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", PORT))
	if err != nil {
		log.Fatalf("Err listening on port %v, %v", PORT, err)
	}

	fmt.Printf("Listening on port:%v", PORT)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &Server{})
	pb.RegisterCustomerServiceServer(grpcServer, &Server{})

	if err = grpcServer.Serve(list); err != nil {
		log.Fatalf("Err serving:%v", err)
	}

}
