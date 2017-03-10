package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	stairs := make([]int, N+1)
	for i := 1; i <= N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &stairs[i])
	}

	// initialize dp table
	dp := make([]int, int(math.Max(float64(N+1), float64(2))))
	dp[0] = 0
	dp[1] = stairs[1]
	if N >= 2 {
		dp[2] = int(math.Max(float64(dp[1]+stairs[2]), float64(stairs[2])))
	}

	for i := 3; i <= N; i++ {
		max := float64(0)
		stair_sum := 0
		for k := 2; k <= 3 && i-k >= 0; k++ {
			stair_sum += stairs[i-k+2]
			max = math.Max(float64(max), float64(dp[i-k]+stair_sum))
		}
		dp[i] = int(max)
	}

	fmt.Printf("%d\n", dp[N])
}
