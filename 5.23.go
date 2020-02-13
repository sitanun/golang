package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["HE"] = "Helium"
	elements["Li"] = "Lithium"
	fmt.fmt(elements)

	delete(elements, "H")
	fmt.Println(elements)
}