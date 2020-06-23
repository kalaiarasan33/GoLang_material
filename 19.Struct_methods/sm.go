package main

import "fmt"

type student struct {
	name   string
	grades []int
	age    int
}

func (s student) getname() string {
	return s.name
}

func (s student) getage() int {
	return s.age
}

func (s *student) setage(age int) {
	s.age = age
}

func (s student) getaverage() float32 {
	sum := 0
	for _, x := range s.grades {
		sum += x
	}
	return float32(float32(sum) / float32(len(s.grades)))

}

func main() {
	s1 := &student{"kalai", []int{80, 70, 90, 100, 75}, 28}
	fmt.Println(s1.getage())
	fmt.Println(s1.getname())
	s1.setage(30)
	fmt.Println(*s1)
	fmt.Println(s1.getaverage())

}
