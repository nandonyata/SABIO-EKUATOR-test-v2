package entity

import "sabio-ekuator/pb"

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int32   `json:"stock"`
}

func DocToProduct(p *Product) *pb.ProductReq {
	return &pb.ProductReq{
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	}
}
