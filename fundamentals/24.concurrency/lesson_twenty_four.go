package lesson_twenty_four

import (
	"fmt"
)

func printTo15() {
	for i := 1; i <= 15; i++ {
		fmt.Println("printTo15() :", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		fmt.Println("printTo10() :", i)
	}
}

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func printTo6(channel1 chan int, channel2 chan int) {
	for i := 1; i <= 3; i++ {
		fmt.Println(<-channel1)
	}
	for i := 1; i <= 3; i++ {
		fmt.Println(<-channel2)
	}
}

func Start() string {
	fmt.Println("Lesson Twenty Four Started...")

	// concurrency
	// - allows us to have multiple blocks of code running at the same time
	// - concurrent tasks in go are called goroutines (not threads)
	// - cannot trust in what order they are going to execute
	/*
		simple example
		go printTo10()
		go printTo15()
		time.Sleep(2 * time.Second)
	*/
	// - channels are a way to communicate between goroutines
	// - channels also help us force goroutines to run in a specific order
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	printTo6(channel1, channel2)

	return "Lesson Twenty Four Complete!"
}
