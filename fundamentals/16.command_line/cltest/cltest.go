package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

// To run this program (exe) in windows:
//
//   - Open powershell
//   - Run: go build ./16.lesson_sixteen/cltest/cltest.go
//	 - A cltest.exe should now be within 16.lesson_sixteen module
//   - Run: ./cltest.exe.
func main() {
	pl(os.Args)

	// can make a console app game with just
	// the executables...nice

	// FindMaxValue()
	// PrintAllValues()
}

// Running the command ./cltest.exe with options:
//   - ./cltest 1 2 3 89 33 29292
//   - Max value : 29292 <--- should be in output
func FindMaxValue() {
	pl("cltest FindMaxValue() start")

	args := os.Args[1:] // get the first index
	var iArgs = []int{}
	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	// get the max value from the maximum values passed
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	pl("Max value :", max)

	pl("cltest FindMaxValue() end")
}

// Running the command ./cltest.exe with options:
//  1. ./cltest 3 6 9
//  2. Example output:
//
// 	  Here is the key:  0
//
// 	  Here is the value:  3
// 	  
//	  Here is the key:  1
// 	  
//	  Here is the value:  6
// 	  
// 	  Here is the key:  2
//
// 	  Here is the value:  9
func PrintAllValues() {
	pl("cltest PrintAllValues() start")
	args := os.Args[1:]
	var iArgs = []int{}
	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	for key, val := range iArgs {
		pl("Here is the key: ", key)
		pl("Here is the value: ", val)
	}
	pl("cltest PrintAllValues() end")
}