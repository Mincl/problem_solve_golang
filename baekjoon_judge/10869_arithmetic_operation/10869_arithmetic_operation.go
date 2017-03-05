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

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &A)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &B)

	fmt.Printf("%d\n", A+B)
	fmt.Printf("%d\n", A-B)
	fmt.Printf("%d\n", A*B)
	fmt.Printf("%d\n", A/B)
	fmt.Printf("%d\n", A%B)
}
