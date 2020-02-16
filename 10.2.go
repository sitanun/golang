package main

import (
	"fmt"
	"myMath"
)

func main() {
	avg := myMath.Average(15, 16.2, 20.3, 32)
	fmt.Println(avg)
}