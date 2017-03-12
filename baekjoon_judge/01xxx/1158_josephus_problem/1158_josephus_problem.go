package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var N, M int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    scanner.Scan()
    fmt.Sscan(scanner.Text(), &N)
    scanner.Scan()
    fmt.Sscan(scanner.Text(), &M)

    people := make([]int, N)
    for i := 1; i <= N; i++ {
        people[i-1] = i
    }

    fmt.Printf("<")
    idx := -1
    for i := 0; i < N; i++ {
        idx += M
        idx %= len(people)
        fmt.Printf("%d", people[idx])
        if i != N-1 {
            fmt.Printf(", ")
        }
        people = append(people[0:idx], people[idx+1:len(people)]...)
        idx--
    }
    fmt.Printf(">")
}
