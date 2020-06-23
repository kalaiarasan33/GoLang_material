
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your age: ")
	scanner.Scan()
	ans := scanner.Text()

	ansint, _ := strconv.ParseInt(ans, 10, 64)

	switch ansint {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 10:
		fmt.Println("three to 10")
	default:
		fmt.Println("Its an int value")
	}

}
