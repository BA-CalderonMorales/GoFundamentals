package lesson_eight

import (
	"fmt"
	"time"
)

var log = fmt.Println

func Start() string {
	log("Lesson Eight Started.")
	log("Working with Time")

	innerLessonOne()

	return "Lesson Eight Complete!"
}

func innerLessonOne() {
	now := time.Now()
	log(now.Year(), now.Month(), now.Day())
	log(now.Hour(), now.Minute(), now.Second())
}
