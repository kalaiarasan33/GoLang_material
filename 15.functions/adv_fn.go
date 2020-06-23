package main

import "fmt"

func add(x, y int) int {
	c := x + y
	return c
}

// Passing func to another function

func test2(myfunc func(x, y int) int) {
	fmt.Println(myfunc(10, 20))
}

// function closure

func returnfunc(x string) func() {
	return func() { fmt.Println(x) }
}
func main() {

	// func is also one of the datatype

	add := func(x, y int) int {
		c := x + y
		return c
	}
	fmt.Println(add(10, 20))

	// Not sure how to use in real, but its one of the type

	sub := func(x, y int) int {
		c := x - y
		return c
	}(10, 20)

	fmt.Println(sub)

	func() {

		fmt.Println("Execute function  without any func call")
	}()

	// Passing func to another function
	test2(add)

	// function closure

	returnfunc("Hello")() // pass parameter and trigger return function
}
