package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	staringsA :=time.Now()
	var a string
	for i := 0; i < 10000; i++ {
		a += "a"
	}
fmt.Prinln("a", time.Since(startA))

starB := time.Now()
}