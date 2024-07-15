package main

var firstName string

// in this statement we declared firstName as a string variable
var address, lastName string

// in this statement we declared address and lastName as string variables using comma we can declare
// more than one variable
var age, height int

// in this statement we declared age and height as int variables
var (
	name, faculty string
	new_age       int
)
// here is another method of declaring variables using the var keyword and () parentheses
// with variable name and type

// note: just like in python it is note necessary to declare the type of the variable because the data type
//  will be same as the value assigned to the variable example:
var(
	stdName="Raj"
	stdlastName="Nayabhari"
	stdAge=21
)
// another way to initialize variables are
var(
	stdName1,stdlastName1,stdAge1="Krisha","Gosain",21
)