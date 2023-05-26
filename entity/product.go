package entity

import "fmt"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       int
}

func (prd *Product) PrintDetail() {
	fmt.Println("ID:", prd.Id)
	fmt.Println("Name:", prd.Name)
	fmt.Println("Description:", prd.Description)
	fmt.Println("Price:", prd.Price)
}
