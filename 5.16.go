package main

import "fmt"

func main() {
	humans := []string{"Abc", "Def", "Dhi"}
	names := []string{"Cartoon", "Sitanun"}
	names = append(names, humans...)
	fmt.Println(names)
}