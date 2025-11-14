package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= 2*n-1; i++ {
		var spaces, stars string
		if i < n {
			spaces = strings.Repeat(" ", n-i)
			stars = strings.Repeat("*", i*2-1)
		} else {
			spaces = strings.Repeat(" ", i-n)
			stars = strings.Repeat("*", 4*n-2*i-1)
		}
		fmt.Printf("%s%s\n", spaces, stars)
	}
}
