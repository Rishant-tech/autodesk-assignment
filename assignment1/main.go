package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go run ./assignment1 <text>")
		os.Exit(1)
	}

	fmt.Println(Reverse(os.Args[1]))
}
