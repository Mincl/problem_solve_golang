package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &A)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &B)
		if A == 0 && B == 0 {
			break
		}
		fmt.Printf("%d\n", A+B)
	}
}
