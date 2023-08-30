package lesson_ten

import (
	"fmt"
)

var log = fmt.Println

func Start() string {
	log("Lesson Ten Started.")
	log("Format your text/string output")
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n",
		"Stuff", 1, 'A', 3.14, true, 1, 1)

	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)

	// save formatted text to variable
	sp1 := fmt.Sprintf("%9.f\n", 3.151592)
	log(sp1)
	return "Lesson Ten Complete!"
}
