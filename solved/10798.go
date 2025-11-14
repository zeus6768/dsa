package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	r = 5
	c = 15
)

var (
	io = bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	)
	board = [r][c]rune{}
)

func main() {
	for i := range r {
		for j := range c {
			board[i][j] = -1
		}
	}

	s := ""
	for i := range r {
		fmt.Fscanln(io, &s)
		for j, k := range []rune(s) {
			board[i][j] = k
		}
	}
	result := []rune{}
	for i := range c {
		for j := range r {
			v := board[j][i]
			if v != -1 {
				result = append(result, board[j][i])
			}
		}
	}
	fmt.Fprintln(io, string(result))
	io.Flush()
}
