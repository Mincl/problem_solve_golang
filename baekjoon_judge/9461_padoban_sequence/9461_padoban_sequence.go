package main

import "fmt"

func main() {
	var T int
	var N int
	init := [6]int{0, 1, 1, 1, 2, 2}
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		dp := make([]int, N+1)

		if N <= 5 {
			fmt.Println(init[N])
		} else {
			for j := 1; j <= N; j++ {
				if j <= 5 {
					dp[j] = init[j]
				} else {
					dp[j] = dp[j-1] + dp[j-5]
				}
			}
			fmt.Println(dp[N])
		}
	}
}
