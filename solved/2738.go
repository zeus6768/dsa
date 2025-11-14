package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var io = bufio.NewReadWriter(
	bufio.NewReader(os.Stdin),
	bufio.NewWriter(os.Stdout),
)

func main() {
	defer io.Flush()
	var n, m int
	fmt.Fscanln(io, &n, &m)

	A := make([][]int, n)
	for r := range n {
		A[r] = make([]int, m)
		for c := range m {
			fmt.Fscan(io, &A[r][c])
		}
	}

	B := make([][]int, n)
	for r := range n {
		B[r] = make([]int, m)
		for c := range m {
			fmt.Fscan(io, &B[r][c])
		}
	}

	C := make([][]int, n)
	for r := range n {
		C[r] = make([]int, m)
		for c := range m {
			C[r][c] = A[r][c] + B[r][c]
		}
	}

	for r := range n {
		var row []string
		var output string
		for c := range m {
			row = append(row, strconv.Itoa(C[r][c]))
			output = strings.Join(row, " ")
		}
		fmt.Fprintln(io, output)
	}
}
