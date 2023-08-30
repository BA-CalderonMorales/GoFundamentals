package lesson_thirteen

import (
	"fmt"
)

var pl = fmt.Println

func Start() string {
	pl("Lesson Thirteen Started.")
	pl("Arrays")

	// Collection of values with the same data type
	// The size of the array cannot change
	// the default values for an array are:
	// 0 : int
	// 0.0 : floats
	// false : boolean
	// "" : string

	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("Index 0:", arr2[0])
	pl("Length :", len(arr2))

	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	// multi dimensional arrays
	arr3 := [2][2]int{
		{1, 2},
		{3, 4}, // must end with semi-colon
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d : %d, value is: %d\n",
				i, j, arr3[i][j])
		}
	}

	// creating an array from a string using slice
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

	// convert a byte array into a string via slice [:]
	byteArray := []byte{'a', 'b', 'c'}
	bStr := string(byteArray[:])
	pl("I'm a string", bStr)

	// slices are like arrays but they CAN grow
	// var name []datatype
	sL1 := make([]string, 6)
	sL1[0] = "Society"
	sL1[1] = "of"
	sL1[2] = "the"
	sL1[3] = "Simulated"
	sL1[4] = "Universe"
	pl("Slice Size :", len(sL1))
	for i := 0; i < len(sL1); i++ {
		pl(sL1[i])
	}

	// key, value in range...
	for _, v := range sL1 {
		pl(v)
	}

	sArr := [5]int{1, 2, 3, 4, 5}
	sL2 := sArr[0:2]               // up to but not including 2nd index
	pl("1st 3 values :", sArr[:3]) // slices...0 is not necessary
	pl("Last 3 valuse :", sArr[2:])
	pl("sL2 :", sL2) // slices are mutable and will mutate the original array...

	// for example:
	sL2[0] = 29
	pl("sArr ", sArr)

	sL2 = append(sL2, 89) // add to the end of the slice
	pl("sL2 :", sL2)
	pl("sArr :", sArr)

	sl3 := make([]string, 6)
	pl("sl3 :", sl3)
	pl("sl3[0] :", sl3[0])

	return "Lesson Thirteen Complete!"
}
