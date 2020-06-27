package main

import "fmt"

func main() {

	go run(1)
	go run(2)
	go run(3)
	fmt.Println("done")

}

func run(x int) {
	fmt.Println(x)
}
