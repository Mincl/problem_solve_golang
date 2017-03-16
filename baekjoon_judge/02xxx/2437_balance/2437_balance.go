package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type ByVal []int

func (v ByVal) Len() int           { return len(v) }
func (v ByVal) Less(i, j int) bool { return v[i] < v[j] }
func (v ByVal) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	sinkers := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sinkers[i])
	}
	sort.Sort(ByVal(sinkers))
	minSeq := 0
	result := 0
	for i := 0; i < N; i++ {
		if sinkers[i] > minSeq+1 {
			result = minSeq + 1
			break
		} else {
			result = minSeq + sinkers[i] + 1
			minSeq = minSeq + sinkers[i]
		}
	}
	fmt.Printf("%d\n", result)
}
