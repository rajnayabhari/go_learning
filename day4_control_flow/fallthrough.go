package main
import "fmt"
func main(){
	switch num:=49;{
	case num<50:
		fmt.Println("number is less than 50")
		fallthrough
	case num<100:
		fmt.Println("number is less than 100")
		fallthrough
	case num<200:
		fmt.Println("number is less than 200")
		
	}
}
// hence in switch case if one condition is true others are not checked
// but using fallthrough we can check all the cases