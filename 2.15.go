package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println.(strings.Trim("-Hello-world", "-"))
	fmt.Println(strings.Trim("+Hello World-", "+"))
}