package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args[1:] {
		fmt.Println(arg)
	}
}
