package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func divide(square [][]int, cutCnt, stI, edI, stJ, edJ int) int {
	if cutCnt == 2 {
		cnt := 0
		for i := stI; i <= edI; i++ {
			for j := stJ; j <= edJ; j++ {
				cnt += square[i][j]
			}
		}
		return cnt
	}

	max := 0
	// vertical cut
	for j := stJ; j < edJ; j++ {
		rCut := divide(square, 2, stI, edI, stJ, j) * divide(square, cutCnt+1, stI, edI, j+1, edJ)
		lCut := divide(square, cutCnt+1, stI, edI, stJ, j) * divide(square, 2, stI, edI, j+1, edJ)
		max = int(math.Max(float64(max), math.Max(float64(rCut), float64(lCut))))
	}

	// horizontal cut
	for i := stI; i < edI; i++ {
		dCut := divide(square, 2, stI, i, stJ, edJ) * divide(square, cutCnt+1, i+1, edI, stJ, edJ)
		uCut := divide(square, cutCnt+1, stI, i, stJ, edJ) * divide(square, 2, i+1, edI, stJ, edJ)
		max = int(math.Max(float64(max), math.Max(float64(uCut), float64(dCut))))
	}
	return max
}

func main() {
	var N, M int
	var mapText string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N, &M)
	square := make([][]int, N)
	for i := 0; i < N; i++ {
		square[i] = make([]int, M)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &mapText)
		for j := 0; j < len(mapText); j++ {
			square[i][j] = int(mapText[j] - '0')
		}
	}
	fmt.Printf("%d\n", divide(square, 0, 0, N-1, 0, M-1))
}
