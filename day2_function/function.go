package main
import(
	"fmt"
	"os"
	"strconv"
)

// func main(){
// 	number1,_:=strconv.Atoi(os.Args[1])
	// If you need to access command-line arguments in Go, you can do it with the os package and the os.
	//Args variable, which holds all the arguments passed to the program.
	// go run function.go 5 4 (here go run function.go runs the program and passes 5 and 4 as
	//  arguments for main function)
// 	number2,_:=strconv.Atoi(os.Args[2])
// 	fmt.Println("Sum is:",number1+number2)
// }

// Syntax for creating custom function
// func name(parameters) (results) {
//     body-content
// }

// Example of custom  function for sum and multiplication
func calc(number1 string, number2 string) (sum int, mul int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    sum = int1 + int2
    mul = int1 * int2
    return
}

func main() {
    sum, mul := calc(os.Args[1], os.Args[2])
    // sum, _ := calc(os.Args[1], os.Args[2])
    // use this code to only print sum and omit mul
    fmt.Println("Sum:", sum)
    fmt.Println("Mul:", mul)
    // to print only sum we need to remove print mul line
}