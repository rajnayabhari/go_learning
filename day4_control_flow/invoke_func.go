package main
import (
"fmt"
"time"
)
func main(){
	switch time.Now().Weekday().String(){
		case "Sunday","Monday","Tuesday","Wednesday","Thursday","Friday":
			fmt.Println("Its time to learn lets go")
		default:
			fmt.Println("Its the weekend,time to rest")
	}
	fmt.Println(time.Now().Weekday().String())
}

// When you call a function from a switch statement,
//  you can modify its logic without changing the expression because you always
//   validate what the function returns.