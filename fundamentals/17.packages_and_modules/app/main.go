package main

import (
	"fmt"
	stuff "lesson_seventeen/my_package"
	"reflect"
)

// package_name "module_name/package_name"
// in order to get stuff.Name to work within this
// project, I had to add the go.work file
// vs code provided an intellisense option
// i selected the option to add the list of
// packages to the workspace (Fundamentals)
// see go.work at the root of this project.

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStringArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}
