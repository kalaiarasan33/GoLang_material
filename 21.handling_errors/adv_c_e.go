// Detailed  error
package main

import (
	"errors"
	"fmt"
	"os"
)

type Requesterror struct {
	StatusCode int
	Err        error
}

func (r *Requesterror) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func Dorequest() (string, error) {
	return "", &Requesterror{
		StatusCode: 503,
		Err:        errors.New("Service unavialble"),
	}

}

func main() {
	s, err := Dorequest()
	if err != nil {
		fmt.Println("Custom error : ", err)
		os.Exit(1)
	}
	fmt.Println("The string:", s)

}
