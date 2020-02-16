package main

import (
	"fmt"
)

type myError struct {
	error string
}

func (e myError) Error() string {
	return e.error
}
func say(ord string) error {
	if word == "hi" {
		return myError{"can't say hi"}
	}
	return nil
}
func main() {
	e1 := say("hello")
}