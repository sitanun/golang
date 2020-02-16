package main

import "fmt"

func hello(t interface{}) {
	fmt.Printf("Hello %T \n", t)
}
func hi(t ...interface{}) {
	fmt.Print("Hi ")
	for _, v := range t {
		fmt.Printf(" %T", v)
	}
	fmt.Println()
}

func main() {
	hello(3.451878655458)
	hello(true)
}