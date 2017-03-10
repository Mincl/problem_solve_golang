package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N >= 90 {
		fmt.Print("A")
	} else if N >= 80 {
		fmt.Print("B")
	} else if N >= 70 {
		fmt.Print("C")
	} else if N >= 60 {
		fmt.Print("D")
	} else {
		fmt.Print("F")
	}
}
