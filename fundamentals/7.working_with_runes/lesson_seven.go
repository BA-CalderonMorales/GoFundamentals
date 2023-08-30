package lesson_seven

import (
	"fmt"
	"unicode/utf8"
)

var log = fmt.Println
var logF = fmt.Printf

func Start() string {
	log("Lesson Seven Started.")
	log("Working with Runes")

	runes()
	return "Lesson Seven Complete!"
}

func runes() {
	rStr := "abcdefg"
	log("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		// i is the index
		// runeVal is the individual value
		// range is each of the values
		logF("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
