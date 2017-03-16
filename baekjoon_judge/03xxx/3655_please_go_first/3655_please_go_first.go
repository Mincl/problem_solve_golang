package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, N int
	var seq string
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &T)
	for i := 0; i < T; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &N)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &seq)
		groupCnt := make(map[byte]int)
		groupLastIdx := make(map[byte]int)
		prevSum := 0
		for j := N - 1; j >= 0; j-- {
			if _, ok := groupLastIdx[seq[j]]; ok {
				prevSum += (groupLastIdx[seq[j]] - j) * 5
			} else {
				groupLastIdx[seq[j]] = j
			}
			groupCnt[seq[j]]++
		}

		sortedSum := 0
		for _, val := range groupCnt {
			sortedSum += 5 * val * (val - 1) / 2
		}
		fmt.Printf("%d\n", prevSum-sortedSum)
	}
}
