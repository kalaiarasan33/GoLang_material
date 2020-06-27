package main

import (
	"fmt"
	"os"
)

type Myerror struct{}

func (M *Myerror) Error() string {
	return "boom"
}

func Sayhello() (string, error) {
	return "", &Myerror{}

}
func main() {
	s, err := Sayhello()
	if err != nil {
		fmt.Println("Custom error : ", err)
		os.Exit(1)
	}
	fmt.Println("The string:", s)

}
