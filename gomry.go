package main

import (
	"fmt"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("gomry: %w", err))
	}
	fmt.Println("path: ", path)
	bob(path)
}
