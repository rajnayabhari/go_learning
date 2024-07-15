package main

import (
	"fmt"
)

func main() {
	firstName := "Raj"
	updateName(&firstName)
	fmt.Println(firstName)
}
func updateName(name *string) {
	*name = "Krisha"
}

// The & operator takes the address of the object that follows it.
// The * operator dereferences a pointer.
//  It gives you access to the object at the address contained in the pointer.
