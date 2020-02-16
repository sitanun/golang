package main

import (
	"fmt"
	"io/ioutil"
)
func main() {
	bs, err := ioutil.Readfil("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}