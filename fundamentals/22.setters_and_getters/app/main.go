package main

import (
	"fmt"
	utils "lesson_twenty_two/my_package"
	"log"
)

func main() {
	fmt.Println("Hello", utils.Name)

	date := utils.Date{}
	err := date.SetDay(21) // set variables using setters
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12) // set variables using setters
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(2018) // set variables using setters
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("1st Day : %d/%d/%d\n",
		date.Day(), date.Month(), date.Year()) // access via getters
}
