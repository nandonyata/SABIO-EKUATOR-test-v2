package entity

import "sabio-ekuator/pb"

type Order struct {
	Id       int32    `json:"id"`
	Quantity int32    `json:"quantity"`
	Total    float64  `json:"total"`
	Customer Customer `json:"customer"`
	Product  Product  `json:"product"`
}

func DocToOrder(o *Order) *pb.Order {
	return &pb.Order{
		Id:         o.Id,
		Quantity:   o.Quantity,
		Total:      o.Total,
		CustomerId: DocToCustomer(&o.Customer),
		ProductId:  DocToProduct(&o.Product),
	}
}
