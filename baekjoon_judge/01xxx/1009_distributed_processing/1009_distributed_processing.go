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
	for i := 0; i < T; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &a, &b)

		res := 1
		for j := 0; j < b; j++ {
			res *= a
			res %= 10
		}
		if res == 0 {
			res = 10
		}
		fmt.Printf("%d\n", res)
	}
}
