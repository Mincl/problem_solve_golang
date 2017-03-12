package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

// Node -
type Node struct {
    x, y int
}

// Length - return length vector selfs
func (n Node) Length() float64 {
    return math.Sqrt(math.Pow(float64(n.x), 2) + math.Pow(float64(n.y), 2))
}

// Add - return additional vector
func (n Node) Add(n2 Node) Node {
    return Node{
        x: n.x + n2.x, y: n.y + n2.y,
    }
}

// Sub - return substract vector
func (n Node) Sub(n2 Node) Node {
    return Node{
        x: n.x - n2.x, y: n.y - n2.y,
    }
}

func backTracking(nodes []Node, N, idx, cnt int, nodeSum Node, nodeSum2 Node) {
    nodeSum2 = nodeSum2.Add(nodes[idx])
    if cnt >= N/2 {
        nodeSum3 := nodeSum.Sub(nodeSum2)
        res := nodeSum2.Sub(nodeSum3)
        //fmt.Printf("<%d,%d> - <%d,%d> = %.6f\n", nodeSum2.x, nodeSum2.y, nodeSum3.x, nodeSum3.y, res.Length())
        min = math.Min(min, res.Length())
        return
    }
    for j := idx + 1; j < N; j++ {
        backTracking(nodes, N, j, cnt+1, nodeSum, nodeSum2)
    }
}

var min float64

func main() {
    var T int
    var N int

    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    fmt.Sscan(scanner.Text(), &T)

    for t := 0; t < T; t++ {
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &N)

        node := make([]Node, N)
        nodeSum := Node{0, 0}
        for i := 0; i < N; i++ {
            scanner.Scan()
            fmt.Sscan(scanner.Text(), &node[i].x, &node[i].y)
            nodeSum = nodeSum.Add(node[i])
        }

        min = 1 << 32
        for i := 0; i < N/2; i++ {
            backTracking(node, N, i, 1, nodeSum, Node{0, 0})
        }

        fmt.Printf("%.6f\n", min)
    }
}
