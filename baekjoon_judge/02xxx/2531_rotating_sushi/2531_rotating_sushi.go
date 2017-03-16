package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var N, d, k, c int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N, &d, &k, &c)
	sushi := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sushi[i])
	}

	cnt := make([]int, d+1)
	maxCnt := 0
	curCnt := 0
	leftIdx := 0
	rightIdx := 0
	for leftIdx < N {
		cnt[sushi[rightIdx]]++
		if cnt[sushi[rightIdx]] == 1 {
			curCnt++
		}
		if (leftIdx <= rightIdx && rightIdx-leftIdx >= k) ||
			(rightIdx < leftIdx && N-leftIdx+rightIdx >= k) {
			cnt[sushi[leftIdx]]--
			if cnt[sushi[leftIdx]] == 0 {
				curCnt--
			}
			leftIdx++
		}
		tmp := curCnt
		if cnt[c] == 0 {
			tmp++
		}
		maxCnt = int(math.Max(float64(maxCnt), float64(tmp)))
		rightIdx++
		if rightIdx >= N {
			rightIdx = 0
		}
	}
	fmt.Printf("%d\n", maxCnt)
}
