package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Fscan(io, &n)
	count := 0
	for range n {
		word := ""
		fmt.Fscan(io, &word)
		if isGroupWord(word) {
			count++
		}
	}
	fmt.Println(count)
}

func isGroupWord(s string) bool {
	exists := make(map[rune]bool)
	for c := 'a'; c < 'A'; c++ {
		exists[c] = false
	}
	var prev rune
	for _, c := range s {
		if prev == c {
			continue
		}
		if exists[c] {
			return false
		}
		exists[c] = true
		prev = c
	}
	return true
}
