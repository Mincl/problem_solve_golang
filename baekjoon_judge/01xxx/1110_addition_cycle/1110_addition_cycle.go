package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)

    cur := N
    for cnt := 1; ; cnt++ {
        cur = (cur%10)*10 + (cur/10+cur%10)%10
        //fmt.Printf("[%d] = %d\n", cnt, cur)
        if cur == N {
            fmt.Print(cnt)
            break
        }
    }
}
