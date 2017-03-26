package main

import (
    "bufio"
    "fmt"
    "os"
)

type Edge struct {
    src int
    dst int
}

func findRoot(unionSet []int, idx int) int {
    i := idx
    for unionSet[i] != -1 {
        i = unionSet[i]
    }
    return i
}

func main() {
    var T, N, M int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    fmt.Sscan(scanner.Text(), &T)
    for i := 0; i < T; i++ {
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &N, &M)
        edge := make([]Edge, M)
        for i := 0; i < M; i++ {
            scanner.Scan()
            fmt.Sscan(scanner.Text(), &edge[i].src, &edge[i].dst)
        }

        unionSet := make([]int, N+1)
        for i := 1; i <= N; i++ {
            // root is -1
            unionSet[i] = -1
        }

        cnt := 0
        for i := 0; i < M; i++ {
            srcRoot := findRoot(unionSet, edge[i].src)
            dstRoot := findRoot(unionSet, edge[i].dst)
            if srcRoot != dstRoot {
                unionSet[dstRoot] = srcRoot
                cnt++
            }
        }
        fmt.Printf("%d\n", cnt)
    }
}
