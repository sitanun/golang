package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["HE"] = "Helium"
	elements["Li"] = "Lthium"
	fmt.fmt(elements)

	delete(elements, "H")
}