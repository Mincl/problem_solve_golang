package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type ByLen []string

func (l ByLen) Len() int      { return len(l) }
func (l ByLen) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l ByLen) Less(i, j int) bool {
	if len(l[i]) == len(l[j]) {
		return strings.Compare(l[i], l[j]) < 0
	}
	return len(l[i]) < len(l[j])
}

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	words := make([]string, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		words[i] = scanner.Text()
	}
	sort.Sort(ByLen(words))
	prev := ""
	for i := 0; i < N; i++ {
		if prev != words[i] {
			fmt.Printf("%s\n", words[i])
			prev = words[i]
		}
	}
}
