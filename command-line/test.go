package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	if len(os.Args) == 1 {
		fmt.Println(os.Args[0])
		return
	}

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}