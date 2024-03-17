package main

import (
	"context"
	"database/sql"
	"fmt"
	"sabio-ekuator/config"
	"sabio-ekuator/entity"
	"sabio-ekuator/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOrder(_ context.Context, req *pb.Order) (*pb.OrderMessage, error) {

	if req.ProductId == nil {
		return nil, status.Error(codes.Aborted, "productId required")
	}
	if req.CustomerId == nil {
		return nil, status.Error(codes.InvalidArgument, "customerId required")
	}

	var customerID int
	row := config.DB.QueryRow(`
		SELECT id FROM Customer
		WHERE id = $1
	`, req.CustomerId.Id)

	if err := row.Scan(&customerID); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Customer not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	var product entity.Product
	rowProduct := config.DB.QueryRow(`
		SELECT id, price, stock FROM Product
		WHERE id = $1
	`, req.ProductId.Id)
	if err := rowProduct.Scan(&product.Id, &product.Price, &product.Stock); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Product not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	resultStock := product.Stock - req.Quantity
	if resultStock < 0 {
		return nil, status.Error(codes.Aborted, "Stock not available")
	}

	totalPrice := req.Quantity * int32(product.Price)
	_, err := config.DB.Exec(`
		INSERT INTO "order" (customer_id, product_id, quantity, total)
		VALUES ($1, $2, $3, $4)
	`, customerID, product.Id, req.Quantity, totalPrice)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	_, err = config.DB.Exec(`
		UPDATE Product
		SET stock = $1
		WHERE id = $2
	`, resultStock, product.Id)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.OrderMessage{
		Message: "Success create order",
	}, nil
}

func (s *Server) FetchOneOrder(_ context.Context, req *pb.OrderId) (*pb.Order, error) {

	order := entity.Order{}
	query := `
		SELECT 
			o.id, 
			o.quantity, 
			o.total,
			c.id as customer_id,
            c.name as customer_name,
            c.email as customer_email,
			p.id as product_id,
			p.name as product_name,
			p.price as product_price,
			p.stock as product_stock
		FROM 
			"order" o
		JOIN
			Customer c ON o.customer_id = c.id
		JOIN
			Product p ON o.product_id = p.id
		WHERE 
			o.id = $1
	`

	row := config.DB.QueryRow(query, req.Id)
	err := row.Scan(&order.Id, &order.Quantity, &order.Total, &order.Customer.Id, &order.Customer.Name, &order.Customer.Email, &order.Product.Id, &order.Product.Name, &order.Product.Price, &order.Product.Stock)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return entity.DocToOrder(&order), nil
}

func (s *Server) FetchAllOrder(_ *pb.OrderEmpty, stream pb.OrderService_FetchAllOrderServer) error {
	query := `
		SELECT 
			o.id, 
			o.quantity, 
			o.total,
			c.id as customer_id,
            c.name as customer_name,
            c.email as customer_email,
			p.id as product_id,
			p.name as product_name,
			p.price as product_price,
			p.stock as product_stock
		FROM 
			"order" o
		JOIN
			Customer c ON o.customer_id = c.id
		JOIN
			Product p ON o.product_id = p.id
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for rows.Next() {
		order := entity.Order{}

		err := rows.Scan(&order.Id, &order.Quantity, &order.Total, &order.Customer.Id, &order.Customer.Name, &order.Customer.Email, &order.Product.Id, &order.Product.Name, &order.Product.Price, &order.Product.Stock)

		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("error scanning: %v\n", err))
		}

		stream.Send(entity.DocToOrder(&order))
	}

	return nil
}
