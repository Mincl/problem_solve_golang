package main

import (
    "bufio"
    "fmt"
    "math"
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
    woods := make([]int, N)
    max := 0
    for i := 0; i < N; i++ {
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &woods[i])
        max = int(math.Max(float64(max), float64(woods[i])))
    }

    st := 0
    ed := max
    max = 0
    for {
        if ed < st {
            break
        }
        height := (ed + st) / 2
        //fmt.Printf("(%d %d %d)\n", st, height, ed)
        sum := 0
        for i := 0; i < N; i++ {
            sum += int(math.Max(float64(0), float64(woods[i]-height)))
        }

        if sum >= M {
            max = int(math.Max(float64(max), float64(height)))
            st = height + 1
        } else {
            ed = height - 1
        }
    }

    writer := bufio.NewWriter(os.Stdout)
    fmt.Fprintf(writer, "%d\n", max)
    writer.Flush()
}
