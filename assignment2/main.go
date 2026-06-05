package main

import (
	"fmt"
	"os"
)

func main() {
	sourcePath := os.Getenv("SourcePath")
	buildNum := os.Getenv("BuildNum")

	if err := UpdateBuildFiles(sourcePath, buildNum); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("build files updated successfully")
}
