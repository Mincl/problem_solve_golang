package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

type ByVal []int

func (v ByVal) Len() int           { return len(v) }
func (v ByVal) Less(i, j int) bool { return v[i] > v[j] }
func (v ByVal) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

func main() {
    var N int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    fmt.Sscan(scanner.Text(), &N)
    theory := make([]int, N)
    for i := 0; i < N; i++ {
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &theory[i])
    }
    sort.Sort(ByVal(theory))
    for i := 10000; i >= 0; i-- {
        cnt := 0
        for j := 0; j < N; j++ {
            if theory[j] >= i {
                cnt++
            }
        }
        if cnt >= i {
            fmt.Printf("%d\n", i)
            break
        }
    }
}
