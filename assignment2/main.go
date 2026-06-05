package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: go run ./assignment2 <value> <delta>")
		os.Exit(1)
	}

	value, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "value must be an integer")
		os.Exit(1)
	}

	delta, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "delta must be an integer")
		os.Exit(1)
	}

	fmt.Println(UpdateValue(value, delta))
}
