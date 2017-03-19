package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Asc []int
type Dsc []int

func (a Asc) Len() int           { return len(a) }
func (a Asc) Less(i, j int) bool { return a[i] < a[j] }
func (a Asc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (d Dsc) Len() int           { return len(d) }
func (d Dsc) Less(i, j int) bool { return d[i] > d[j] }
func (d Dsc) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &A[i])
	}
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &B[i])
	}
	sort.Sort(Asc(A))
	sort.Sort(Dsc(B))
	sum := 0
	for i := 0; i < N; i++ {
		sum += A[i] * B[i]
	}
	fmt.Printf("%d\n", sum)
}
