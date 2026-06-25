package main

import (
	"os"
	"fmt"
)

func main() {
	var s string
	var sep string

	if len(os.Args) == 1 {
		fmt.Println(os.Args[0])
		return
	}

	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
	}
	
	fmt.Println(s)
}