package main

import "fmt"

func findPosition(tallerCnt, position, check []int, N int, cnt int) bool {
	if cnt == N {
		for i := 0; i < N; i++ {
			taller := 0
			for j := 0; j < i; j++ {
				if position[j] > position[i] {
					taller++
				}
			}
			if taller != tallerCnt[position[i]-1] {
				return false
			}
		}
		return true
	} else {
		for i := 0; i < N; i++ {
			if check[i] == 0 {
				check[i] = 1
				position[cnt] = i + 1
				if findPosition(tallerCnt, position, check, N, cnt+1) {
					return true
				}
				check[i] = 0
			}
		}
		return false
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	tallerCnt := make([]int, N)
	position := make([]int, N)
	check := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tallerCnt[i])
	}
	if findPosition(tallerCnt, position, check, N, 0) {
		for i := 0; i < N; i++ {
			fmt.Printf("%d ", position[i])
		}
	}
}
