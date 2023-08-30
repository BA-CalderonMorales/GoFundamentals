package lesson_fifteen

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func Start() string {
	pl("Lesson Fifteen Started.")
	pl("File I/O")

	// f - file
	// err - error
	f, err := os.Create("15.lesson_fifteen/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // whenever out of scope, it closes the file
	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	// convert integer array to string array values
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	// write the newly created string array values to the file, handle errors
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	// open the created file
	f, err = os.Open("15.lesson_fifteen/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// read from the file and print out all the lines one by one
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime :", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	// append to a file
	/*
		Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must
		be specified.

		O_RDONLY : open the file read-only
		O_WRONLY : open the file write-only
		O_RDWR   : open the file read-write

		These can be or'ed

		O_APPEND : append data to the file when writing
		O_CREATE : create a new file if none exists
		O_EXCL   : used with O_CREATE, file must not exist
		O_SYNC   : open for synchronous I/O
		O_TRUNC  : truncate regular writable file when opened
	*/

	err = nil // resets err if it's nil from above use...

	// check if the file exists
	_, err = os.Stat("15.lesson_fifteen/data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exist.")
	} else {
		f, err = os.OpenFile(
			"15.lesson_fifteen/data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err = f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}

	return "Lesson Fifteen Complete!"
}
