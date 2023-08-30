package lesson_twenty_seven

import (
	"fmt"
)

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func Start() string {
	fmt.Println("Lesson Twenty Seven Started...")

	fmt.Println("Factorial of 4 :", factorial(4))
	// 1st : result = 4 * factorial(3) = 4 * 6 = 24
	// 2nd : result = 3 * factorial(2) = 3 * 2 = 6
	// 3rd : result = 2 * factorial(1) = 2 * 1 = 2
	// 4th : result = 1 * factorial(0) = 1 * 1 = 1

	return "Lesson Twenty Seven Complete!"
}
