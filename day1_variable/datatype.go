package main

import "fmt"

// integer
var integer8 int8 = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807

// float]
var float1 float32 = 2147483647
var float2 float64 = 9223372036854775807

// boolean
var featureFlag bool = true
var featureFlag2 bool = false

// string
var name1 string = "Raj"
var name2 string = "Nayabhari"

// most common example of using escape characters
// \n for new lines
// \r for carriage returns
// \t for tabs
// \' for single quotation marks
// \" for double quotation marks
// \\ for backslashes
func main(){
	fullName := "Raj Nayabhari \t(Bca \"ktm\")\n"
	fmt.Println(fullName)
}