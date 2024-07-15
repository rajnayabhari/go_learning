package main
import "fmt"
func location(city string)(string,string){
	var region string
	var continent string
	switch city{
	case "mumbai","delhi","chennai","kochi","hyderabad":
		region,continent="india","asia"
	case"kathmandu","lalitpur","lumbini","bhaktapur":
		region,continent="nepal","asia"
	case"irvine","san diego","los angeles","san jose":
		region,continent="us","north america"
	default:
		region,continent="unknown","unknown"
	}
	return region ,continent
}
func main(){
	region,continent:=location("kathmandu")
	fmt.Printf("Raj works in %s,%s\n",region,continent)

}