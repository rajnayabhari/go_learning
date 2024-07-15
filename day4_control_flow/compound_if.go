package main
import "fmt"
func givemeannumber() int{
	return 10
}
func main(){
	if num :=givemeannumber();num<0{
		fmt.Println("Negative number")
	}else if num< 10{
		fmt.Println("number has only one digit")
	}else{
		fmt.Println("number has more than one digit")
	}
	// fmt.Println(num)
	// this line creates error because num is not defined 
	// In Go, declaring variables within if blocks is idiomatic.
	// It's a way to program effectively by using a convention that's common in Go.
}