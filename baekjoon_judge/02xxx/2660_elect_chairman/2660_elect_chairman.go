package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func floydWarshall(member [][]int, N int) {
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if member[i][k] > 0 && member[k][j] > 0 &&
					(member[i][j] == 0 ||
						member[i][j] > member[i][k]+member[k][j]) {
					member[i][j] = member[i][k] + member[k][j]
				}
			}
		}
	}
}

func main() {
	var N, st, ed int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	member := make([][]int, N)
	for i := 0; i < N; i++ {
		member[i] = make([]int, N)
	}
	for {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &st, &ed)
		if st == -1 && ed == -1 {
			break
		}
		member[st-1][ed-1] = 1
		member[ed-1][st-1] = 1
	}
	floydWarshall(member, N)
	min := 1 << 30
	for i := 0; i < N; i++ {
		max := 0
		for j := 0; j < N; j++ {
			if i != j {
				max = int(math.Max(float64(max), float64(member[i][j])))
			}
		}
		min = int(math.Min(float64(min), float64(max)))
	}

	list := make([]int, N)
	cnt := 0
	for i := 0; i < N; i++ {
		max := 0
		for j := 0; j < N; j++ {
			if i != j {
				max = int(math.Max(float64(max), float64(member[i][j])))
			}
		}
		if max == min {
			list[cnt] = i
			cnt++
		}
	}

	fmt.Printf("%d %d\n", min, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Printf("%d ", list[i]+1)
	}
}
