package lesson_six

import (
	"fmt"
	"strings"
)

var log = fmt.Println

func Start() string {
	log("Lesson Six Started.")
	log("More string operations")

	innerLessonOne()

	return "Lesson Six Complete!"
}

func innerLessonOne() {
	sV1 := "A word"
	// strings are just arrays of bytes...
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	log(sV2)
	log("Length :", len(sV2))
	log("Contains Another:", strings.Contains(sV2, "Another"))
	log("o index :", strings.Index(sV2, "o"))
	log("Replace :", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words\n"      // \t \" \' \\ escaped chars
	sV3 = strings.TrimSpace(sV3) // gets rid of whitespace
	log("Split :", strings.Split("a-b-c-d", "-"))
	log("Lower :", strings.ToLower(sV2))
	log("Upper :", strings.ToUpper(sV2))
	log("Prefix :", strings.HasPrefix("tacocat", "taco"))
	log("Suffix :", strings.HasSuffix("tacocat", "cat"))
}
