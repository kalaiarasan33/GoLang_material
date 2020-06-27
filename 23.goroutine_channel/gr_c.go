package main

import "fmt"

func main() {
	mychannel := make(chan int)

	go run(1, mychannel)
	// get the value from channel and storing in variable
	text1 := <-mychannel
	// print value from variable
	fmt.Println(text1)

	go run(2, mychannel)
	text2 := <-mychannel
	fmt.Println(text2)
	go run(3, mychannel)
	text3 := <-mychannel
	fmt.Println(text3)
	fmt.Println("done")

}

// function is doing some operation and sending value to channel
func run(x int, c chan int) {
	x += 10
	c <- x
}
