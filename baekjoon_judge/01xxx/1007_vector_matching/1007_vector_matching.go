package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Node -
type Node struct {
	num, x, y int
}

// Vector -
type Vector struct {
	stN, edN, vx, vy int
}

// Length - return length vector selfs
func (v Vector) Length() float64 {
	return math.Sqrt(math.Pow(float64(v.vx), 2) + math.Pow(float64(v.vy), 2))
}

func getVector(a, b Node) Vector {
	return Vector{
		stN: a.num, edN: b.num,
		vx: b.x - a.x, vy: b.y - a.y,
	}
}

func main() {
	var T int
	var N, M int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &T)

	for t := 0; t < T; t++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &N)

		node := make([]Node, N)
		for i := 0; i < N; i++ {
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &node[i].x)
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &node[i].y)
			node[i].num = i
		}

		// generate vector
		M = 0
		vector := make([]Vector, N*(N-1))
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if i != j {
					vector[M] = getVector(node[i], node[j])
					M++
				}
			}
		}
		//fmt.Println(vector)

		// allocate DP tables
		cap := N/2 + 1
		d := make([][]float64, cap)
		dVectorSum := make([][]Vector, cap)
		dPointSet := make([][][]int, cap)
		for i := 0; i < cap; i++ {
			d[i] = make([]float64, M)
			dVectorSum[i] = make([]Vector, M)
			dPointSet[i] = make([][]int, M)
			for j := 0; j < M; j++ {
				dPointSet[i][j] = make([]int, N)
			}
		}

		// initialize DP tables
		for i := 0; i < M; i++ {
			dVectorSum[1][i] = vector[i]
			d[1][i] = vector[i].Length()
			dPointSet[1][i][vector[i].stN] = 1
			dPointSet[1][i][vector[i].edN] = 1
		}

		// run DP
		for i := 2; i < cap; i++ {
			for j := 0; j < M; j++ {
				minLen := float64(1 << 30)
				var minVx, minVy int
				var minVec Vector
				for k := 1; k <= j; k++ {
					if d[i-1][j-k] != -1 && dPointSet[i-1][j-k][vector[j].stN] != 1 && dPointSet[i-1][j-k][vector[j].edN] != 1 {
						vx := dVectorSum[i-1][j-k].vx + vector[j].vx
						vy := dVectorSum[i-1][j-k].vy + vector[j].vy
						len := math.Sqrt(math.Pow(float64(vx), 2) + math.Pow(float64(vy), 2))
						if minLen > len {
							minVx = vx
							minVy = vy
							minVec = vector[j]
							minLen = len
						}
					} else if d[i][j-k] != -1 {
						len := math.Sqrt(math.Pow(float64(dVectorSum[i][j-k].vx), 2) + math.Pow(float64(dVectorSum[i][j-k].vy), 2))
						if minLen > len {
							minVx = dVectorSum[i][j-k].vx
							minVy = dVectorSum[i][j-k].vy
							minVec = dVectorSum[i][j-k]
							minLen = len
						}
					}
				}
				if minLen < float64(1<<29) {
					d[i][j] = minLen
					dVectorSum[i][j] = Vector{
						stN: minVec.stN, edN: minVec.edN, vx: minVx, vy: minVy,
					}
					dPointSet[i][j][minVec.stN] = 1
					dPointSet[i][j][minVec.edN] = 1
				} else {
					d[i][j] = -1
				}
			}
		}

		// for i := 0; i < M; i++ {
		// 	fmt.Println(vector[i])
		// }
		//
		// for i := 0; i < cap; i++ {
		// 	for j := 0; j < M; j++ {
		// 		fmt.Printf("%5.2f ", d[i][j])
		// 	}
		// 	fmt.Printf("\n")
		// }
		//print result
		fmt.Printf("%.6f\n", d[cap-1][M-1])
	}
}
