package main

import "fmt"

func man() {
	elements := make(map[string]map[string]string)
	elements["S"] = map[string]string{"name": "Hydrogen", "state": "gas"}
	elements["she"] = map[string]string{"name": "Hydrogen", "state": "gas"}
	fmt.Println(elements)
	fmt.Println(elements["she"]["state"])
}