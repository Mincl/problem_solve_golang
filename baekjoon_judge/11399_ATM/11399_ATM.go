package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type ByVal []int

func (v ByVal) Len() int           { return len(v) }
func (v ByVal) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByVal) Less(i, j int) bool { return v[i] < v[j] }

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	people := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &people[i])
	}
	sort.Sort(ByVal(people))
	waitTime := 0
	sum := 0
	for i := 0; i < N; i++ {
		waitTime += people[i]
		sum += waitTime
	}
	fmt.Printf("%d\n", sum)
}
