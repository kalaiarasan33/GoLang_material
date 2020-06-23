package main

import "fmt"

func main() {

	// var arrs []string

	// fmt.Println("Empty String Array", arrs)

	// var arri  []int

	// fmt.Println("Empty integer Array", arri)

	var arrs [5]string

	fmt.Println("Empty String Array", arrs)

	//Empty String Array [    ]

	var arri [5]int

	fmt.Println("Empty integer Array", arri)

	//Empty integer Array [0 0 0 0 0]

	arrint := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Integer Array", arrint)
	// changed 0 index value
	arrint[0] = 11
	fmt.Println("Integer Array", arrint)
	fmt.Println("Length Array", len(arrint))

	ar := [5]int{100, 200, 300, 400, 500}
	sum := 0

	for i := 0; i <= len(ar)-1; i++ {
		fmt.Println("Index value iteration", ar[i])
		sum += ar[i]
	}
	fmt.Println("Iterated sum value: ", sum)

	arr2D := [2][2]int{{1, 2}, {2, 3}}
	fmt.Println("Two Dimension Array", arr2D)
	fmt.Println("Two Dimension Array 0 index array ", arr2D[0])
	fmt.Println("Two Dimension Array 0 index array, 1 element value", arr2D[0][1])

}
