package main

import "fmt"

type num struct {
	x int
	y int
}

// if you use pointer it will change value in object with help of reference memoryid
func change(p *num) {
	p.x = 10
}

// if you use value, it will just  copy a value to canonical variable
func copy(p num) {
	p.x = 10
	fmt.Println(p)
}

type circle struct {
	radius float64
	center *num
}
type seecircle struct {
	radius float64
	*num
}

func main() {
	var p num = num{y: 20}

	fmt.Println(p)
	change(&p) // using pointer
	fmt.Println(p)
	fmt.Println("-------------------------------------------------------------------------")
	pp := num{y: 30}
	fmt.Println(pp)
	copy(pp)
	fmt.Println(pp)
	// ###########################################################################################################

	// using struct in another struct
	fmt.Println("-------------------------------------------------------------------------")
	cn := &num{30, 40}
	c1 := circle{10.56, cn} // c1 := circle{10.56, &num{30, 40}} another method to define
	fmt.Println(c1)
	fmt.Println(c1.center)
	fmt.Println(c1.center.x)

	// using struct in another struct
	fmt.Println("-------------------------------------------------------------------------")

	c11 := seecircle{20.56, &num{40, 20}}
	fmt.Println(c11)
	fmt.Println(c11.x)

}
