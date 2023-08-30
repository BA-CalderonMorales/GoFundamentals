package lesson_twenty

import (
	"fmt"
)

type Customer struct {
	name    string
	address string
	balance float64
}

func GetCustInfo(c Customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.balance)
}

func NewCustAddress(c *Customer, address string) {
	c.address = address
}
