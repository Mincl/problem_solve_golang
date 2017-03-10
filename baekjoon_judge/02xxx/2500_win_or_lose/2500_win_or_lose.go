package main

import (
	"fmt"
	"math"
)

func q_mul(q []int, q2 []int) []int {
	size := len(q)
	res := make([]int, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i+j >= size {
				break
			}
			res[i+j] += (q[i] % 42043 * q2[j] % 42043) % 42043
		}
	}

	return res
}

func main() {
	var N, K, T int
	var panel_val int
	fmt.Scanf("%d %d %d\n", &N, &K, &T)
	panel_cnt := make([]int, T)
	for i := 0; i < T; i++ {
		fmt.Scanf("%d\n", &panel_val)
		panel_cnt[panel_val]++
	}

	q := make([]int, K+1)
	q2 := make([]int, K+1)
	q_res := make([]int, K+1)

	// initialize
	for idx, v := range panel_cnt {
		if idx > K {
			break
		}
		q[idx] = v
	}

	for i := N; i > 0; {
		log := int(math.Floor(math.Log2(float64(i))))
		copy(q2, q)
		for j := 0; j < log; j++ {
			q2 = q_mul(q2, q2)
		}
		if i == N {
			copy(q_res, q2)
		} else {
			q_res = q_mul(q_res, q2)
		}
		i -= int(math.Pow(2, float64(log)))
	}

	sum := 0
	for j := 0; j < K+1; j++ {
		sum += q_res[j]
	}
	fmt.Printf("%d", sum%42043)
}
