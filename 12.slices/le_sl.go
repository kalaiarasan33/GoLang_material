package main

import "fmt"

func main() {

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[2:3])

	arrr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrr[2:4])

	var s []int = arrr[:]
	fmt.Println(s)

	ss := make([]int, 0)
	fmt.Printf("%T", ss)

	ss = arrr[2:3]
	fmt.Printf("\n%v", ss)

}
