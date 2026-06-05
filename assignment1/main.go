package main

import (
	"fmt"
	"os"
)

func main() {
	testStr := "String; 2be reversed..."
	if got := ReverseWords(testStr); got != "gnirtS; eb2 desrever..." {
		panic(fmt.Sprintf("reverse_words(%q) = %q", testStr, got))
	}

	if len(os.Args) > 1 {
		fmt.Println(ReverseWords(os.Args[1]))
	}
}
