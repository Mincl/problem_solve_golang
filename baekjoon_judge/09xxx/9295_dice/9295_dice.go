package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	var a, b int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &T)
	for i := 1; i <= T; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &a, &b)
		fmt.Printf("Case %d: %d\n", i, a+b)
	}
}
