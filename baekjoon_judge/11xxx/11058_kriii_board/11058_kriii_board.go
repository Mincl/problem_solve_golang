package main

import (
    "fmt"
    "math"
)

func main() {
    var N int
    fmt.Scan(&N)

    d := make([]int, N+1)
    d[0] = 0
    for i := 1; i <= N; i++ {
        d[i] = int(math.Max(float64(d[i]), float64(d[i-1]+1)))
        for j := 1; 2+j <= i; j++ {
            d[i] = int(math.Max(float64(d[i]), float64(d[i-(2+j)]*(j+1))))
        }
    }
    fmt.Printf("%d\n", d[N])
}
