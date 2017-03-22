package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var N, M, K int
	var a, b, c int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N, &M, &K)
	path := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		path[i] = make([]int, N+1)
	}
	for i := 0; i < K; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &a, &b, &c)
		if a < b && path[a][b] < c {
			path[a][b] = c
		}
	}

	d := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		d[i] = make([]int, N+1)
	}
	for i := 1; i <= N; i++ {
		if path[1][i] > 0 {
			d[2][i] = path[1][i]
		}
	}

	for i := 3; i <= M; i++ {
		for j := 1; j <= N; j++ {
			max := 0
			for k := 1; k <= N; k++ {
				if path[k][j] > 0 && d[i-1][k] > 0 {
					max = int(math.Max(float64(max), float64(d[i-1][k]+path[k][j])))
				}
			}
			d[i][j] = max
		}
	}

	max := 0
	for i := 2; i <= M; i++ {
		max = int(math.Max(float64(max), float64(d[i][N])))
	}
	fmt.Printf("%d\n", max)
}
