package lesson_twenty

import (
	"fmt"
)

type Contact struct {
	fName, lName, phone string
}

type Business struct {
	name, address string
	Contact       // composition over inheritance
}

func (b Business) Info() {
	fmt.Printf("Conact at %s is %s %s",
		b.name, b.Contact.fName, b.Contact.lName)
}
