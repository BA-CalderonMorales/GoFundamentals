package lesson_twelve

import (
	"fmt"
)

var pl = fmt.Println

func Start() string {
	pl("Lesson Twelve Started.")
	pl("Using range")

	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}

	return "Lesson Twelve Complete!"
}
