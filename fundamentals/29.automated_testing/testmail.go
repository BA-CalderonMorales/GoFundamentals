package lesson_twenty_nine

import (
	"fmt"
	"regexp"
)

func PrintMail() {
	fmt.Println("Hello World")
}

func IsEmail(s string) (string, error) {
	re, _ := regexp.Compile(`[\w.%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`)
	if re.MatchString(s) {
		return "Valid Email", nil // valid email, no error
	} else {
		return "", fmt.Errorf("Not a valid email")
	}
}
