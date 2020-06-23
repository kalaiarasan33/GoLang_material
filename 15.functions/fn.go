package main

import "fmt"

func add(x, y int) int {
	c := x + y
	return c
}

func swapparams(x, y int) (int, int) {
	a := x
	b := y
	return b, a
}

func swap2params(x, y int) (b int, a int) {
	a = x
	b = y
	return
}
func main() {
	a := add(10, 20)
	fmt.Println("Sum value: ", a)
	b, a := swapparams(10, 20)
	fmt.Println("swap value: ", b, a)
	u, v := swap2params(1, 2)
	fmt.Println("swap2 value: ", u, v)

}
