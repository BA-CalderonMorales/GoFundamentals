package lesson_twenty_three

import (
	"fmt"
)

func Start() string {
	fmt.Println("Lesson Twenty Three Started...")
	// interfaces = contracts -> if anything inherits it, the methods must be implemented

	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()

	fmt.Println("Cats Name :", kitty2.Name())
	return "Lesson Twenty Three Complete!"
}
