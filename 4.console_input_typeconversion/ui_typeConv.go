package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Your name  is %T %v", input, input)

	// type comversion example

	fmt.Println("\nEnter your age: ")
	scanner.Scan()
	//   strconv.ParseInt(scanner.Text(), 10, 64) func return output,err also if it not able operate
	// so creating two varible may be _ or err
	intinput,_ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Your name  is %T %v", intinput, intinput)
	

}
