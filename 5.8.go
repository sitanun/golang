package main

import "fmt"

func main() {
	alphabets := [2][3]string{{"a", "b", "c"}, {"A", "B", "C"}}
	fmt.Prinln(alphabets)
	fmt.Prinln(alphabets[0][1])
	numbers := [2][3][2]int{
		{
			{1, 2},
			{10, 20},
			{100, 200},
		},
		{
			{8, 9},
			{80, 99},
			{800, 900},
    	},
    }
	fmt.Prinln(numbers)
	fmt.Prinln(numbers[1][2][0])
}