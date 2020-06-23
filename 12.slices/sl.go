package main

import "fmt"

func main() {

	//[0 1 2 3 4]
	// pointer - memory location id of index
	// length - length size of array
	// capacity - maximum  lenght size of array

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// initializing slice
	var s []int = arr[:]
	fmt.Println("Initialized slice", s)

	// slicing 2,3 element in arr index
	var se []int = arr[1:3]
	fmt.Println("sliced element", se)
	fmt.Println("length of slice", len(se)) // 2,3
	fmt.Println("cap of slice", cap(se))    // 2, 3, 4, 5
	fmt.Println(se[:cap(se)])
	fmt.Println("slicing from sliced arr", se[1:])

	var arrsl []int = []int{1, 2, 3, 4, 3, 4, 5, 6, 7, 3, 2, 33} // it will initialize array with slice, will setup predefined things
	fmt.Println(arrsl[2:4])
	fmt.Println("capacity", cap(arrsl[2:4]))
	c := append(arrsl, 10)
	fmt.Println("append value", c)
	// we can initialize slice
	a := make([]int, 0) //special built-in function that is used to initialize slices, maps, and channels
	// a := make([]int, 0,10) increase capcity size
	a = arrsl[1:9]
	fmt.Println("make slice", a)
}
