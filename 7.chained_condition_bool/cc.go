package main

import "fmt"

func main() {
	//	!  Not
	//	|| OR
	//	&& AND
	val := (5 > 6) || (2 < 1)
	fmt.Printf("Dtype %T - value: %v", val, val)
	oval := (true || false) && !false
	fmt.Printf("\nDtype %T - value: %v", oval, oval)

}

