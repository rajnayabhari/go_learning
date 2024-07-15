package main

import "fmt"
func main(){
	var integer64 int=99999999999999999999999999999
	//  this statement causes error because it causes overflow for int type
	fmt.Println(integer64)
}