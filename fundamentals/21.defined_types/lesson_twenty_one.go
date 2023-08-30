package lesson_twenty_one

import (
	"fmt"
)

func Start() string {
	fmt.Println("Lesson Twenty One Started.")

	ml1 := Millileter(TeaSpoon(3) * 4.92)
	fmt.Printf("3 teaspoons = %.2f\n", ml1)

	ml2 := Millileter(TableSpoon(3) * 14.79)
	fmt.Printf("3 tablespoons = %.2f\n", ml2)

	fmt.Println("2 teaspoons + 4 teaspoons =", TeaSpoon(2) + TeaSpoon(4))

	fmt.Println("2 teaspoons > 4 teaspoons =", TeaSpoon(2) > TeaSpoon(4))

	// Ameatur conversions...
	fmt.Printf("3 teaspoons = %.2f millileters\n", TeaSpoonToMillileter(3))
	fmt.Printf("3 tablespoons = %.2f millileters\n", TableSpoonToMillileter(3))

	// Using implied function, sort of like overriding in c#
	tsp1 := TeaSpoon(3)
	fmt.Printf("%.2f teaspoons = %.2f millileters\n", tsp1, tsp1.ToMillileters())
	tbsp1 := TableSpoon(3)
	fmt.Printf("%.2f tablespoons = %.2f millileters\n", tbsp1, tbsp1.ToMillileters())

	return "Lesson Twenty One Complete!"
}
