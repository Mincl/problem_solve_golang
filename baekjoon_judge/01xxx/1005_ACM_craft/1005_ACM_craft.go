package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var T int
	var N, K, W int
	var ST, ED int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &T)
	for i := 0; i < T; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &N)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &K)
		graph := make([][]int, N)
		building := make([]int, N)
		indegree := make([]int, N)
		required_time := make([]int, N)
		for j := 0; j < N; j++ {
			graph[j] = make([]int, 0)
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &building[j])
		}

		for j := 0; j < K; j++ {
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &ST)
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &ED)
			graph[ST-1] = append(graph[ST-1], ED-1)
			indegree[ED-1]++
		}
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &W)

		zero_indegree_set_size := 0
		zero_indegree_set := make([]int, N)
		for idx, v := range indegree {
			if v == 0 {
				zero_indegree_set[zero_indegree_set_size] = idx
				zero_indegree_set_size++
			}
		}

		for j := 0; j < zero_indegree_set_size; j++ {
			idx := zero_indegree_set[j]
			required_time[idx] += building[idx]
			for _, idx2 := range graph[idx] {
				required_time[idx2] = int(math.Max(float64(required_time[idx2]), float64(required_time[idx])))
				indegree[idx2]--
				if indegree[idx2] == 0 {
					zero_indegree_set[zero_indegree_set_size] = idx2
					zero_indegree_set_size++
				}
			}
		}

		fmt.Println(required_time[W-1])
	}
}
