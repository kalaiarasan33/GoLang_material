package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Read file
	f, err := ioutil.ReadFile("filehandling/test.txt")
	if err != nil {
		fmt.Println("Error in filereading", err)
		return
	}
	fmt.Println(string(f))
	// Write file
	m := []byte("Hello kalai")
	ioutil.WriteFile("filehandling/kalai.txt", m, 777)

	of, err := os.Create("filehandling/oskalai.txt")
	defer of.Close()

	if err != nil {
		fmt.Println("os file handling err")
	}
	of.WriteString("Hello os module")

	of.Sync()

	w := bufio.NewWriter(of)
	w.WriteString("writing in buffio")
	w.Flush()
}
