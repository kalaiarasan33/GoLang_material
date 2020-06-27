package main

import (
	"errors"
	"fmt"
	"strings"
)

func Cap(s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty name")
	}
	return strings.ToUpper(s), nil

}

func main() {
	o, err := Cap("")
	if err != nil {
		fmt.Println("Error is : ", err)
	}
	fmt.Println(o)
}
