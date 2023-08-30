package lesson_five

import (
	"fmt"
)

var log = fmt.Println

func Start() string {
	log("Lesson Five Started.")
	log("Operators - e.g. Conditional/logical operators")
	conditionalOperators()
	logicalOperators()
	return "Lesson Five Complete!"
}

func conditionalOperators() {
	// include the following: > < >= <= == !=
	iAge := 8

	if (iAge >= 1) && (iAge <= 18) {
		log("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		log("Important Birthday")
	} else if iAge >= 65 {
		log("Important Birthday")
	} else {
		log("Not an Important Birthday")
	}

	log("!true =", !true)
}

func logicalOperators() {
	// include the following: && || !
}
