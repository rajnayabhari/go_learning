package main

import (
	"fmt"
	"regexp"
)

func main() {
	// regular expression
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)
	contact := "food@bar.com"
	switch {
	case email.MatchString(contact):
		fmt.Println("Its an email")
	case phone.MatchString(contact):
		fmt.Println("Its a Phone number")
	default:
		fmt.Println(contact, "is not recognized")

	}
}
