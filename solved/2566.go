package main

import (
	"bufio"
	"fmt"
	"os"
)

const size = 9

func main() {
	io := bufio.NewReader(os.Stdin)
	var max, value, r, c int
	for i := range size {
		for j := range size {
			fmt.Fscan(io, &value)
			if max < value {
				max = value
				r, c = i, j
			}
		}
	}
	fmt.Printf("%d\n%d %d\n", max, r+1, c+1)
}
