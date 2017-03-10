package main

import (
	"fmt"
)

func comb(n, k int) uint64 {
	return fact(n) / fact(n-k) / fact(k)
}

func fact(n int) uint64 {
	res := uint64(1)
	for i := 1; i <= n; i++ {
		res *= uint64(i)
	}
	return res
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	dp := make([][]uint64, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]uint64, N+1)
	}
	dp[1][1] = 1
	for i := 2; i <= N; i++ {
		dp[i][1] = fact(i - 1)
		dp[i][i] = uint64(1)
		for j := 2; j < i; j++ {
			sum := uint64(0)
			for k := 1; k <= i/2; k++ {
				sum += comb(i, k) * dp[k][1] * dp[i-k][j-1]
			}
			dp[i][j] = sum
		}
	}
	for i := 1; i <= N; i++ {
		sum := uint64(0)
		for j := 1; j <= i; j++ {
			sum += dp[i][j]
		}
		sum2 := uint64(0)
		for j := 1; j <= i; j++ {
			sum2 += dp[i][j]
			fmt.Printf("%4d/%-4d ", sum2, sum)
		}
		fmt.Printf("\n")
	}
}

/*
nCk = n개중 k개를 선택하는 경우의 수

n개의 노드로 하나의 사이클을 만드는 경우의 수 =
    nCn * n-1C1 * n-2C1 * ... * 1C1 = (n-1)!

n개의 노드로 두개의 사이클을 만드는 경우의 수 =
k개와 k'개로 분리 (k + k' = n)
    [1 <= k <= n/2]
    nCk * (k-1)!
    k'Ck' * (k'-1)!
    sum[ nCk*(k-1)!*k'Ck'*(k'-1)! ]

C(n, c)
C(n, 1) = (n-1)!

n = 3
C(3, 1) = 3C3*(3-1)! = 2!
C(3, 2) = 3C1*(1-1)!*C(2,1)
        = 3C1*(1-1)!*2C2*(2-1)! = 3
        = 3C1*(1-1)!(2-1)!
C(3, 3) = 1


n = 4
C(4, 1) = 4C4*(4-1)! = 3!
C(4, 2) = 4C1*C(1,1)*C(3,1) + 4C2*C(2,1)*C(2,1)
        = 4C1*(1-1)!*3C3*(3-1)! + 4C2*(2-1)!*2C2*(2-1)!
        = 8 + 6 = 14
        = 4C1*(1-1)!(3-1)! + 4C2*(2-1)!*(2-1)!
        =
C(4, 3) = 4C1*C(1,1)*C(3,2) + 4C2*C(2,1)*C(2,2)
        = 4*3 + 6 = 18
C(4, 4) = 1

R(4) = C(4, 1) + C(4, 2) + C(4, 3) + C(4, 4)

C(5, 1) = 4!
C(5, 2) = 5C1*C(1,1)*C(4,1) + 5C2*C(2,1)*C(3,1) =
C(5, 3) = 5C1*C(1,1)*C(4,2) + 5C2*C(2,1)*C(3,2) = 6*5!/(3*2)
C(5, 4) = 5C1*C(1,1)*C(4,3) + 5C2*C(2,1)*C(3,3) = 3*5!/(2*2)

C(n, c) =
    [1 <= k <= n/2]
    sum[ nCk*C(k,1)*C(n-k, c-1) ]


*/
