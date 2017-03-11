package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := N; i >= 1; i-- {
		fmt.Printf(fmt.Sprintf("%%%ds\n", N), strings.Repeat("*", i))
	}
}
