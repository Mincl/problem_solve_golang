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
	num := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &num[i])
	}
	sort.Sort(ByVal(num))

	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < N; i++ {
		fmt.Fprintf(writer, "%d\n", num[i])
	}
	writer.Flush()
}
