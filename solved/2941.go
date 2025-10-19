package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	croatianAlphabets := []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}
	reader := bufio.NewReader(os.Stdin)
	var input string
	fmt.Fscan(reader, &input)
	for _, croatianAlphabet := range croatianAlphabets {
		input = strings.ReplaceAll(input, croatianAlphabet, "a")
	}
	fmt.Println(len(input))
}
