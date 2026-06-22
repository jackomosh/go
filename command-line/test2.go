package main

import (
	"os"
	"fmt"
)

func main() {
	var s, sep string

	if len(os.Args) == 1 {
		fmt.Println(os.Args[0])
	}
	
	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
	}
	fmt.Println(s)
}