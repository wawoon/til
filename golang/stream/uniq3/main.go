package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	args := os.Args
	if len(args[1:]) == 0 {
		countLength(os.Stdin, counts)
	} else {
		for _, path := range args[1:] {
			fmt.Print(path)
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("error: %v", err)
			}
			countLength(file, counts)
		}
	}

	for key, val := range counts {
		fmt.Printf("%v: %v\n", key, val)
	}
}

func countLength(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
