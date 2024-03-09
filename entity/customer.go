package entity

import "sabio-ekuator/pb"

type Customer struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func DocToCustomer(c *Customer) *pb.Customer {
	return &pb.Customer{
		Id:    c.Id,
		Name:  c.Name,
		Email: c.Email,
	}
}
