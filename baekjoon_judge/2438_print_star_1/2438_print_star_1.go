package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		fmt.Printf("%s\n", strings.Repeat("*", i))
	}
}
