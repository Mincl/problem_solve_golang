package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

func main() {
    var K int
    fmt.Scan(&K)

    N := 0
    for int(math.Pow(2, float64(N))-2) < K {
        N++
    }
    N--

    lower := math.Pow(2, float64(N)) - 2
    idx := int64(K-int(lower)) - 1
    //fmt.Printf("%d %d %d %d\n", K, int(N), int(lower), idx)
    str := fmt.Sprintf(fmt.Sprintf("%%0%ds\n", N), strconv.FormatInt(idx, 2))
    str = strings.Replace(str, "0", "4", len(str))
    str = strings.Replace(str, "1", "7", len(str))
    fmt.Printf("%s\n", str)
    /*
       2^(N-1) - 2 < K
       K <= 2^N - 2
       K+2 <= 2^N
       Ceil(log2(K+2)) == N
    */
}
