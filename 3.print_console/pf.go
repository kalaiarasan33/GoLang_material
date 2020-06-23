package main

import "fmt"

func main() {

	var name string = "kalai"
	var age int = 27
	var male bool = true
	var interest float64 = 99.9
	// printf is printformation %T display type %v display value, %t display bool, so on check document
	fmt.Printf("\n%T", name)
	fmt.Printf("\n%v", age)
	fmt.Printf("\n%t", male)
	fmt.Printf("\n%f", interest)

	fmt.Println("\n", name, age, male, interest)
	// Sprintf helps to store print output in varibale
	var out string = fmt.Sprintf("\n Myself %v , i am learning go", name)
	fmt.Println(out)

}
