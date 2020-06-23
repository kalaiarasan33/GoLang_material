package main

import "fmt"

func main() {

	// immutable datatype
	var x int = 5
	y := x 
	y = 7
	fmt.Println("Immutable")
	fmt.Println(x,y)
	
	// mutable datatype
	var as []int = []int{1,2,3,4}
	bs:= as
	bs[0]=10
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("mutable")
	fmt.Println(as,bs) 

	var am map[string]int = map[string]int{"one":1,"two":2}
	bm := am
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("mutable")
	bm["ten"]=10
	fmt.Println(am,bm) 
	bm["ten"]=0
	fmt.Println(am,bm) 
	am["three"]=3
	fmt.Println(am,bm) 

	var ar [3]int = [3]int{1,2,3}
	br := ar
	br[0] = 10
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Immutable")
	fmt.Println(ar,br)	
	


}

