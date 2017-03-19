package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
)

type ByVal []int

func (v ByVal) Len() int           { return len(v) }
func (v ByVal) Less(i, j int) bool { return v[i] < v[j] }
func (v ByVal) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

func isPrime(n int) bool {
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var N, c int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    fmt.Sscan(scanner.Text(), &N)

    candy := make([]int, 1)
    candy[0] = 0
    candyCnt := make(map[int]int)
    maxSum := 0
    zeroCnt := 0
    for i := 0; i < N; i++ {
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &c)
        if c == 0 {
            zeroCnt++
            continue
        }
        maxSum += c
        if val, ok := candyCnt[c]; ok {
            candyCnt[c] = val + 1
        } else {
            candyCnt[c] = 1
            candy = append(candy, c)
        }
    }
    sort.Sort(ByVal(candy))

    d := make([][]int, N+1)
    for i := 0; i <= N; i++ {
        d[i] = make([]int, maxSum+1)
    }
    d[0][0] = 1
    for i := 1; i < len(candy); i++ {
        copy(d[i], d[i-1])
        for k := 1; k <= candyCnt[candy[i]]; k++ {
            for j := 1; j <= maxSum; j++ {
                if j-(candy[i]*k) >= 0 {
                    d[i][j] += d[i-1][j-(candy[i]*k)]
                }
            }
        }
    }
    // for i := 0; i < len(candy); i++ {
    //     fmt.Println(d[i])
    // }

    sum := 0
    for j := 2; j <= maxSum; j++ {
        if isPrime(j) {
            sum += d[len(candy)-1][j] * (zeroCnt + 1)
        }
    }
    fmt.Printf("%d\n", sum)
}
