package main

import "fmt"

func main() {
    var N int
    var input, pattern string
    fmt.Scan(&N)
    for i := 0; i < N; i++ {
        fmt.Scan(&input)
        if i == 0 {
            pattern = input
        } else {
            for j := 0; j < len(input); j++ {
                if pattern[j] != input[j] {
                    pattern = pattern[:j] + "?" + pattern[j+1:]
                }
            }
        }
    }
    fmt.Println(pattern)
}
