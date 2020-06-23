package main

import "fmt"

func main() {
	//	At that point you can drop the semicolons: C's while is spelled for in Go.

	// for true {
	// 	fmt.Println("Infinite loop")
	// }
	i := 0
	for i < 5 {
		fmt.Println("Hello :", i)
		i++ // or i+= 1
	}

	for x := 0; x <= 6; x++ {
		if x == 4 {
			continue
			// dont do any operation  or do any operation for 4 , then continue
		} else if x == 5 {
			break
			// break the loop in specific condition
		}
		fmt.Println("for loop Hello :", x)

	}

}
