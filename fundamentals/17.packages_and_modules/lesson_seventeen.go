package lesson_seventeen

import (
	// internal_packages "deep_dive_packages/internal"
	"fmt"
)

var pl = fmt.Println

func Start() string {
	pl("Lesson Seventeen Started.")

	pl("See: ./app package for more information on lesson.")

	// here we'll get a "Use of internal package not allowed" error
	// internal_packages.MY_PRIVATE_KEY = "new private key"

	return "Lesson Seventeen Complete!"
}