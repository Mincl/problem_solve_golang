package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", N-i), strings.Repeat("*", i))
	}
}
