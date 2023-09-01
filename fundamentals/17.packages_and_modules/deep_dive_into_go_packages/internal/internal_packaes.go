package internal_packages

import (
	"fmt"
)

// Provides another scoping mechanism
// Accessible by parent and its children
// Enables better organization without leaking details

var MY_PRIVATE_KEY = "MY_PRIVATE_KEY"

func Start() {
	fmt.Println("Hello from internal package!")

}