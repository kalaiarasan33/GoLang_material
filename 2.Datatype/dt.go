package main

import "fmt"

func main() {
	//  Explicit menthod to  define  variable

	var name string = "kalai"
	var age int = 27
	var male bool = true
	var interest float64 = 99.9
	fmt.Printf("%T", name)
	fmt.Printf("%T", age)
	fmt.Printf("%T", male)
	fmt.Printf("%T", interest)

	fmt.Println("\n", name, age, male, interest)

	// Implicit menthod to  define  variable

	var Iname = "kalai"
	var Iage = 27
	var Imale = true
	var Iinterest = 99.9

	fmt.Printf("%T", Iname)
	fmt.Printf("%T", Iage)
	fmt.Printf("%T", Imale)
	fmt.Printf("%T", Iinterest)
	fmt.Println("\n", Iname, Iage, Imale, Iinterest)

	// Assignment Expression menthod to  define  variable

	aname := "kalai"
	aage := 27
	amale := true
	ainterest := 99.9

	fmt.Printf("%T", aname)
	fmt.Printf("%T", aage)
	fmt.Printf("%T", amale)
	fmt.Printf("%T", ainterest)
	fmt.Println("\n", aname, aage, amale, ainterest)

}
