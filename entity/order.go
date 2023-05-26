package entity

import "time"

type Order struct {
	Id           int
	ProductId    int
	BuyerEmail   string
	BuyerAddress string
	OrderDate    time.Time
	Product      Product
}
