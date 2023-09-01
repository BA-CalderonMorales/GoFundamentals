package aliasing

import (
	"encoding/json"
	"log"
	// "pluralsight.com/libmanager/json" <-- problem
	system "fmt"
	// the "_" is used with db drivers
	_ "log"
)


type Todo struct {
	ID   int
	Text string
}

func Start() {
	todo := Todo{
		ID:   1,
		Text: "Learn Go",
	}
	// ...
	data, err := json.Marshal(todo)
	if err != nil {
		log.Fatal(err)
	}
	system.Println(string(data)) // example of aliasing...
	// ...
}