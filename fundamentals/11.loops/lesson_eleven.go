package lesson_eleven

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func Start() string {
	pl("Lesson Eleven Started.")
	pl("Loops")

	// for initialization; condition;
	// postStatement {BODY}

	for x := 1; x <= 5; x++ {
		pl(x) // x is scoped to for loop
	}

	for x := 5; x >= 1; x-- {
		pl(x) // x cannot be used outside of scope
	}

	fX := 0
	for fX < 5 { // create a "while" loop...gross lol
		pl(fX)
		fX++
	}

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	for true {
		fmt.Print("Guess a number between 0 and 50 :")
		pl("\nRandom Number is :", randNum)
		// un-comment the following lines to see this loop in action
		// reader := bufio.NewReader(os.Stdin)
		// guess, err := reader.ReadString('\n')
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// guess = strings.TrimSpace(guess)
		// iGuess, err := strconv.Atoi(guess)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		var iGuess = randNum // line was added to avoid stopping the lessons
		if iGuess > randNum {
			pl("Pick a Lower Value")
		} else if iGuess < randNum {
			pl("Pick a Higher Value")
		} else {
			pl("You Guessed it!")
			break // breaks out of while loop
		}
	}

	return "Lesson Eleven Complete!"
}
