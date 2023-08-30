package lesson_fourteen

import (
	"fmt"
)

var log = fmt.Println

func sayHello() {
	log("Hello!")
}

func getSum(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero.")
	} else {
		return x / y, nil
	}
}

func getRunningSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func attemptToChangeValue(f3 int) int {
	f3 += 1
	return f3
}

func changeValue(myPtr *int) {
	*myPtr = 12 // will override the pointers value
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2 // will affect the values since it's modifying by reference
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums)) // casting to float64
	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

func Start() string {
	log("Lesson Fouteen Started.")
	log("Functions")

	// func funcName(parameters) returnType {BODY}

	sayHello()
	log("getSum result :", getSum(12, 32))
	log(getTwo(12))
	log(getQuotient(5, 0))            // compiler needs help catching this error (so we handle)
	log(getRunningSum(1, 2, 3, 4, 5)) // pass an array of n-values to a function

	vArr := []int{1, 2, 3, 4}
	log("Array sum :", getArraySum(vArr)) // pass by value (not reference)

	f3 := 5
	log("f3 before func (pass by value) :", f3)
	attemptToChangeValue(f3)
	log("f3 after func (pass by value) :", f3)

	f4 := 10
	log("f4 before func (pass by reference) :", f4)
	changeValue(&f4) // to pass by reference instead of value use the '&' symbol
	log("f4 after func (pass by reference :", f4)

	// Ptr = Pointer (reference)
	f5 := 20
	var f5Ptr *int = &f5
	log("f5 Address :", f5Ptr)
	log("f5 Value :", *f5Ptr) // * symbol is key here
	*f5Ptr = 35
	log("f5 Value :", *f5Ptr)

	// double each value in the array
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	log(pArr)

	// passing a slice to a function
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))

	return "Lesson Fourteen Complete!"
}
