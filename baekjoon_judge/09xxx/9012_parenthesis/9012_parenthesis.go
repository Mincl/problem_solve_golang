package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	var ps string
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &T)
	for i := 0; i < T; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &ps)

		openCnt := 0
		for j := 0; j < len(ps); j++ {
			if ps[j] == '(' {
				openCnt++
			} else {
				openCnt--
			}
			if openCnt < 0 {
				break
			}
		}
		if openCnt != 0 {
			fmt.Printf("NO\n")
		} else {
			fmt.Printf("YES\n")
		}
	}
}
