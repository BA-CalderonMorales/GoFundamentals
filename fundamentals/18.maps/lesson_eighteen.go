package lesson_eighteen

import (
	"fmt"
)

func Start() string {
	fmt.Println("Lesson Eighteen Started.")

	// key value pairs - a map
	// var myMap map [keyType]value
	var heroes map[string]string
	heroes = make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"

	villians := make(map[string]string) // one-liner
	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{
		1: "Krypto",
		2: "Bat Hound",
	}

	fmt.Printf("Batman is %v\n", heroes["Batman"])
	fmt.Println("Chip :", superPets[3])
	_, ok := superPets[3]
	fmt.Println("Is there a 3rd pet :", ok)

	for key, value := range heroes {
		fmt.Printf("%s is %s\n", key, value)
	}

	delete(heroes, "The Flash")

	return "Lesson Eighteen Complete!"
}
