package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	responds := make(chan string)
	args := os.Args
	if len(args[1:]) > 0 {
		for _, url := range args[1:] {
			go handleURL(responds, url)
		}

		for i := 0; i < len(args[1:])+2; i++ {
			select {
			case msg := <-responds:
				fmt.Println(msg)
			}
		}
	}

}

func handleURL(responds chan string, url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	responds <- string(b)
}
