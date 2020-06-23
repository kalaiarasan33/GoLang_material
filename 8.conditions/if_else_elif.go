package main

import "fmt"

func main() {
	age := 16
	// if condtion{

	// }else if condtion{

	// }else {

	// }
	if age > 22 {
		fmt.Println("you can drive")

	} else if (age > 18) && (age < 22) {
		fmt.Println("Please drive carefully")
	} else {
		fmt.Println("You cannot drive")
	}
}
