package main

import (
	"fmt"
)

func main() {
	data := getSizeOfContribution()
	runCommands(data)
	fmt.Println(data)
}
