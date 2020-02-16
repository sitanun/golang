package main

import "fmt"

type I interface {
	F()
}

func desc(i I) {
	fmt.Printlf("%v, %T \n", i,i)
}

type T1 struct {
	text string
}