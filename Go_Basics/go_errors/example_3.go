package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func someFunction() error {
	return &MyError{Code: 404, Message: "Resource not found"}
}

func someError() error {
	return fmt.Errorf("someError failed: %w", errors.New("an underlying error"))
}

func main() {
	err := someFunction()
	if err != nil {
		fmt.Println("Error:", err)
	}

	err2 := someError()
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
}
