package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cant divide '%d' by zero", a)
	}
	return a / b, nil
}

//  fmt package can be used to add dynamic data to the error, such as an int, string, or another error

/* Errors can be returned as nil, and in fact, it’s the default, or “zero”, value of on error in Go.
This is important since checking if err != nil is the idiomatic way to determine if an error was
encountered (replacing the try/catch statements you may be familiar with in other programming languages)*/
