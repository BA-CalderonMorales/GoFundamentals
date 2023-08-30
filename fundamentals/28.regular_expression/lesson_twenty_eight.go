package lesson_twenty_eight

import (
	"fmt"
	"regexp"
)

func Start() string {
	fmt.Println("Lesson Twenty Eight Started...")

	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?", reStr)
	if match {
		fmt.Println(match)
	} else {
		fmt.Println("No match found!")
	}

	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")
	fmt.Println("Match String :", r.MatchString(reStr2))
	fmt.Println("Find String :", r.FindString(reStr2))
	fmt.Println("Index String :", r.FindStringIndex(reStr2))
	fmt.Println("All String :", r.FindAllString(reStr2, -1))
	fmt.Println("First 2 Strings :", r.FindAllString(reStr2, 2))
	fmt.Println("All Submatch Index :", r.FindAllStringSubmatchIndex(reStr2, -1))
	fmt.Println("Replace All String :", r.ReplaceAllString(reStr2, "Dog"))

	return "Lesson Twenty Eight Complete!"
}
