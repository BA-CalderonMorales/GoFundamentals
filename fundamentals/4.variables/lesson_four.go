package lesson_four

import (
	"fmt"
	"reflect"
	"strconv"
)

var log = fmt.Println

func Start() string {
	log("Lesson Four Started.")
	log("Variables are fun! See the lesson for more details.")
	variables()
	return "Lesson Four Complete!"
}

func variables() {
	basics()
	types()
	casting()
}

func basics() {
	log("Basics section...")
	// var name type
	// - name is allowed to contain letters and numbers
	// - Uppercase means it's allowed outside the package
	// - lowercase means it's private, contained within package
	// - camelCase is the convention
	// - squigglies will appear under your variable if it's unused
	// - var keyword is optional
	var vName string = "Brandon"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.4 // ok for initializing a variable
	v4 = 8.0  // must use "=" sign after initialization

	log(vName)
	log(v1, v2)
	log(v3)
	log(v4)
}

func types() {
	log("Types section...")
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	log(reflect.TypeOf(25))
	log(reflect.TypeOf(3.14))
	log(reflect.TypeOf(true))
	log(reflect.TypeOf("Hello"))
	log(reflect.TypeOf('🦍'))
}

func casting() {
	log("Casting section...")
	cV1 := 1.5
	cV2 := int(cV1)
	log(cV2)

	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3) // ASCII to integer
	log(cV4, err, reflect.TypeOf(cV4))

	cV5 := 50000000
	cV6 := strconv.Itoa(cV5) // Integer to ASCII
	log(cV6, reflect.TypeOf(cV6))

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		log(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	log(cV9, reflect.TypeOf(cV9))
}
