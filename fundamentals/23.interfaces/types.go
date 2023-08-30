package lesson_twenty_three

import (
	"fmt"
)

// Cat - Type

type Cat string

// Cat - Methods

func (c Cat) Attack() {
	fmt.Println("Meow! Hiss!")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	fmt.Println("Hiss!")
}

func (c Cat) HappySound() {
	fmt.Println("Purr!")
}
