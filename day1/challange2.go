package main
import "fmt"
func main(){
	const integer uint=-10
	//  this statement causes error because uint only takes positive values
	fmt.Println(integer)
}