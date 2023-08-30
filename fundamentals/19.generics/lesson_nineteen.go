package lesson_nineteen

import (
	"fmt"
)

// generic type parameter is going to be capitalized
type MyConstraint interface {
	int | float64 // my function will work if the param is an int or float...
}

// pkg.go.dev/golang.org/x/exp/constraints <-- built-in constraints

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func Start() string {
	fmt.Println("Lesson Nineteen Started.")

	// with generics, you can specify the datatype
	// to be used at a later time.

	// used to create functions that can work with multiple different data
	// types

	fmt.Println("5 + 4 = ", getSumGen(5, 4))
	fmt.Println("5.6 + 4.7 = ", getSumGen(5.6, 4.7))
	// fmt.Println("5.6 + 4.7 = ", getSumGen("5.6", 4.7)) will throw "string does not implement MyConstraint" error

	return "Lesson Nineteen Complete!"
}
