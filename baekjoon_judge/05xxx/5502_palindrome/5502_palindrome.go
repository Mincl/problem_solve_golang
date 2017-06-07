package main

import (
	"fmt"
)

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	var N int
	var str string
	fmt.Scan(&N)
	fmt.Scan(&str)

	d := make([][]int, N)
	for i := 0; i < N; i++ {
		d[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		d[i][i] = 0
		if i+1 < N && str[i] == str[i+1] {
			d[i][i+1] = 0
		} else if i+1 < N {
			d[i][i+1] = 1
		}
	}

	for k := 2; k < N; k++ {
		for i := 0; i+k < N; i++ {
			j := i + k
			d[i][j] = Min(d[i+1][j]+1, d[i][j-1]+1)
			if str[i] == str[j] {
				d[i][j] = Min(d[i][j], d[i+1][j-1])
			} else {
				d[i][j] = Min(d[i][j], d[i+1][j-1]+2)
			}
		}
	}

	// for i := 0; i < N; i++ {
	// 	fmt.Println(d[i])
	// }

	fmt.Printf("%d\n", d[0][N-1])
}
