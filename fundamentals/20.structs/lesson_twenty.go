package lesson_twenty

import (
	"fmt"
)

func Start() string {
	fmt.Println("Lesson Twenty Started.")

	// from ./customer.go
	var tS Customer
	tS.name = "Tom Smith"
	tS.address = "5 Main St"
	tS.balance = 234.56
	GetCustInfo(tS)
	NewCustAddress(&tS, "123 Bravo St")
	fmt.Println("Address :", tS.address)

	// struct literal
	sS := Customer{
		"Sally Smith",
		"123 Main St",
		0.0,
	}
	fmt.Println("Name :", sS.name)

	// from ./rectangle.go
	rect1 := Rectangle{10.0, 15.0}
	fmt.Println("Rect Area :", rect1.Area())

	// NOTE: Go does not support inheritance

	// from ./contact.g
	contact1 := Contact{
		"James",
		"Wang",
		"555-1212",
	}
	business1 := Business{
		"ABC Plumbing",
		"234 North St",
		contact1,
	}
	business1.Info()

	return "Lesson Twenty Complete!"
}
