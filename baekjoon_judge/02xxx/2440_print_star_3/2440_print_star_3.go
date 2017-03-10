package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Printf("%s\n", strings.Repeat("*", N-i))
	}
}
