package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, C, W int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N, &C, &W)
	wood := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &wood[i])
	}

	// maxi := 0
	max := int64(0)
	woodCnt := 0
	cutCnt := 0
	sale := int64(0)
	cost := int64(0)
	for i := 1; i <= 10000; i++ {
		sale = int64(0)
		cost = int64(0)
		for j := 0; j < N; j++ {
			if wood[j]%i != 0 {
				cutCnt = wood[j] / i
			} else {
				cutCnt = wood[j]/i - 1
			}
			woodCnt = wood[j] / i
			if int64(woodCnt)*int64(i)*int64(W)-int64(cutCnt)*int64(C) > 0 {
				sale += int64(woodCnt) * int64(i) * int64(W)
				cost += int64(cutCnt) * int64(C)
			}
		}
		// fmt.Println(i, woodCnt, cutCnt, sale, cost, sale-cost)
		if max < sale-cost {
			max = sale - cost
			// 	maxi = i
		}
	}
	// fmt.Println(maxi)
	fmt.Printf("%d\n", max)
}
