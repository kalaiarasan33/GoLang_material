package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 4, 55, 66, 7, 8}

	// _ instead of variable , it wont show an error, its not used
	for _, element := range a {
		fmt.Println("First loop", element)

	}

	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && i != j {
				fmt.Println("Nested loop", element)

			}

		}

	}

}
