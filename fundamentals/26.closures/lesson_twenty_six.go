package lesson_twenty_six

import (
	"fmt"
)

func useFunc(f func(int, int) int, x, y int) {
	fmt.Println("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func Start() string {
	fmt.Println("Lesson Twenty Six Started...")

	intSum := func(x, y int) int {
		return x + y
	}
	fmt.Println("5 + 4 =", intSum(5, 4))

	sampl1 := 1

	changeVar := func() {
		sampl1 += 1
	}

	fmt.Println("sampl1 =", sampl1) // sampl1 = 1

	changeVar()

	fmt.Println("sampl1 =", sampl1) // value will change, sampl1 = 2

	fmt.Println("What is 5 + 8?")
	useFunc(sumVals, 5, 8)

	return "Lesson Twenty Six Complete!"
}
