package main

import (
	"bufio"
	"fmt"
	"os"
)

// pipeとredirectionの違い
// pipeは、あるコマンドの標準出力を他のコマンドの標準入力にする
// redirectionは、あるコマンドの標準出力、標準エラーをファイルに書き出しをするときに使う

func main() {
	dict := make(map[string]int)

	args := os.Args
	if len(args[1:]) > 0 {
		for _, arg := range args[1:] {
			dict[arg]++
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			dict[scanner.Text()]++
		}
	}

	for key, val := range dict {
		fmt.Printf("%v: %v\n", key, val)
	}
}
