package main

import (
	"fmt"
)

func main() {
	var bowl string
	fmt.Scan(&bowl)

	height := 0
	for i := 0; i < len(bowl); i++ {
		if i != 0 && bowl[i-1] == bowl[i] {
			height += 5
		} else {
			height += 10
		}
	}

	fmt.Print(height)
}
