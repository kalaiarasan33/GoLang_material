package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(mp)
	fmt.Println(mp["apple"])
	// delete map key and value
	delete(mp, "apple")

	// get value from map in another way
	val, ok := mp["apple"]
	fmt.Println("Key value is not present in map", val, ok)
	val1, ok1 := mp["orange"]
	fmt.Println("Key value is  present in map", val1, ok1)

	// length of the map

	fmt.Println(len(mp))

	fmt.Println(mp)
	// another way to create map
	maps := make(map[string]int)
	maps["test"] = 1
	maps["test2"] = 2
	fmt.Println("make map", maps)

}
