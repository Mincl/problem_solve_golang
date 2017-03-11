package main

import (
	"bufio"
	"fmt"
	"os"
)

func DFS(N int, idx int, adjacency [][]int, is_worm []int) {
	for i := 0; i < N; i++ {
		if adjacency[idx][i] == 1 && is_worm[i] != 1 {
			is_worm[i] = 1
			DFS(N, i, adjacency, is_worm)
		}
	}
}

func main() {
	var N, M int
	var a, b int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &M)
	adjacency := make([][]int, N)
	for i := 0; i < N; i++ {
		adjacency[i] = make([]int, N)
	}
	for i := 0; i < M; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &a, &b)
		adjacency[a-1][b-1] = 1
		adjacency[b-1][a-1] = 1
	}

	is_worm := make([]int, N)
	is_worm[0] = 1
	DFS(N, 0, adjacency, is_worm)
	cnt := 0
	for i := 0; i < N; i++ {
		if is_worm[i] == 1 {
			cnt++
		}
	}
	fmt.Println(cnt - 1)
}
