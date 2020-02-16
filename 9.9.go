package main

import "fmt"

tpe I interface {
	F()
}

type T struct {
	text string
}

func (t T) F() {
	fmt.Println(t.text)
}

func man() {
	var i I
}