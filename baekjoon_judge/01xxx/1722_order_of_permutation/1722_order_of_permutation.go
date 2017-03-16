package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func factorial(n int) int64 {
	f := int64(1)
	for i := 1; i <= n; i++ {
		f *= int64(i)
	}
	return f
}

func kPermutation(k int64, n int, remain []int) {
	if k == 1 || n == 1 {
		for i := 0; i < len(remain); i++ {
			fmt.Printf("%d ", remain[i])
		}
	} else {
		f := factorial(n - 1)
		idx := int(math.Ceil(float64(k)/float64(f))) - 1
		fmt.Printf("%d ", remain[idx])
		kPermutation(k-int64(idx)*f, n-1, append(remain[:idx], remain[idx+1:]...))
	}
}

func main() {
	var N, t int
	var k int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &t)
	remain := make([]int, N)
	for i := 0; i < N; i++ {
		remain[i] = i + 1
	}

	if t == 1 {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &k)
		kPermutation(k, N, remain)
	} else {
		sum := int64(0)
		perm := make([]int, N)
		for i := 0; i < N; i++ {
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &perm[i])
		}
		for i := 0; i < N; i++ {
			for idx, val := range remain {
				if val == perm[i] {
					remain = append(remain[:idx], remain[idx+1:]...)
					sum += int64(idx) * factorial(N-i-1)
				}
			}
		}
		fmt.Printf("%d\n", sum+1)
	}
}
