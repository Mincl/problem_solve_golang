package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	var A, B int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &T)
	for i := 0; i < T; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &A)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &B)
		fmt.Printf("Case #%d: %d + %d = %d\n", i+1, A, B, A+B)
	}
}
