package main

// typical import statement
// import "fmt"

// import block
import (
	aliasing "deep_dive_packages/aliasing"
	internal_packages "deep_dive_packages/internal"
	vendor_directories "deep_dive_packages/vendor_directories"
	"fmt"
)

func main() {
	fmt.Println("Hello, gophers!")

	// alternative import methods

	// aliases 
	// helpful if you have two packages with the same name
	aliasing.Start()

	// import for side effects

	// internal packages
	// can be used here because it's at the same package level
	internal_packages.Start()
	internal_packages.MY_PRIVATE_KEY = "new private key" // is also possible

	// relative imports
	// import package relative to current one	
	// not valid in workspaces or modules
	// used to stub something out relatively quickly
	// modules help resolve this (makes you don't really need to use relative imports)

	// working with vendor directories
	vendor_directories.Start()
}