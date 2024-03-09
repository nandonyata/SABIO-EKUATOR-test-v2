package entity

import (
	"fmt"
	"sabio-ekuator/pb"
)

type Product struct {
	Id    int32   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int32   `json:"stock"`
}

func DocToProduct(p *Product) *pb.ProductReq {
	fmt.Println(p, "masuk ke sini")

	return &pb.ProductReq{
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
		Id:    p.Id,
	}
}
