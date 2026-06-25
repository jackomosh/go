package main

import (
	"os"
	"fmt"
)

func main() {
	// var s, sep string

	if len(os.Args) == 1 {
		fmt.Println(os.Args[0])
	}
	
	for index, arg := range os.Args[1:] {
		fmt.Printf("Index: %d, Argument: %s\n", index+1, arg)
	}
}