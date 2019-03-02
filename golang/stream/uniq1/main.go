package main

import (
	"fmt"
	"os"
)

func main() {
	dict := make(map[string]int)

	args := os.Args
	for _, arg := range args[1:] {
		dict[arg] += 1
	}

	for key, val := range dict {
		fmt.Printf("%v: %v\n", key, val)
	}
}
