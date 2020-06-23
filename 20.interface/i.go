package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getarea(s shape) float64 {
	return s.area()

}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	shapes := []shape{r, c} // Create slice with shape  as data type to use interface
	shapes[0].area()
	shapes[1].area()
	// using for loop to use interface
	for _, sh := range shapes {
		fmt.Println(sh.area())
		// use method to call interface
		fmt.Println("method: ", getarea(sh))
	}

}
