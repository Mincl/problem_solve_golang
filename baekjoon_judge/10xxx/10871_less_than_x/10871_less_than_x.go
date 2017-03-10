package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, X int
	var a int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &X)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &a)
		if a < X {
			fmt.Printf("%d ", a)
		}
	}
}
