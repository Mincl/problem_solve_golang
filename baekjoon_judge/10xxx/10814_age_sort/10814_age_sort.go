package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Member struct {
	RegNum int
	Age    int
	Name   string
}

type ByAge []Member

func (a ByAge) Len() int      { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool {
	if a[i].Age == a[j].Age {
		return a[i].RegNum < a[j].RegNum
	}
	return a[i].Age < a[j].Age
}

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	members := make([]Member, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &members[i].Age)
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &members[i].Name)
		members[i].RegNum = i
	}

	sort.Sort(ByAge(members))
	for i := 0; i < N; i++ {
		fmt.Printf("%d %s\n", members[i].Age, members[i].Name)
	}
}
