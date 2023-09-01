package vendor_directories

import (
	"fmt"

	"github.com/google/uuid"
)

func Start() {
	fmt.Println("Hello from vendor directories!")
	// apply to workspaces only
	// module system does not support vendor directories
	// first official version management strategy
	// hierarchical directory structure, resolved at compile time

	uid := uuid.New()
	// this will still print out the same thing since
	// modules do not support vendor directories
	fmt.Println(uid)
	
	// to overrite what .New() returns,
	// we go to the package that's being imported in
	// the current file, in this case, github.com/google/uuid
	// copy and paste that package (uuid) into ./vendor/github.com/google
	// if you want to be evil, just create the package name and a src.go file
	// the src.go file should theoretically contain all the functions you need
	// now just call that function in the same way you would uuid.New()
	// this is just a crazy way to do method overriding imo
}