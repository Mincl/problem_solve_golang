package main

import (
    "bufio"
    "fmt"
    "os"
)

type Data struct {
    pos int
    cnt int
}

type Queue struct {
    data []Data
}

func (q *Queue) Push(d Data) {
    q.data = append(q.data, d)
}

func (q *Queue) Pop() Data {
    d := q.data[0]
    q.data = q.data[1:]
    return d
}

func (q *Queue) Size() int {
    return len(q.data)
}

func main() {
    var N, K int
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    fmt.Sscan(scanner.Text(), &N, &K)
    dup := make([]bool, 500001)
    q := new(Queue)
    q.Push(Data{N, 0})
    dup[N] = true
    for q.Size() > 0 {
        d := q.Pop()
        if d.pos == K {
            fmt.Printf("%d\n", d.cnt)
            return
        }
        if 0 <= d.pos-1 && !dup[d.pos-1] {
            q.Push(Data{d.pos - 1, d.cnt + 1})
            dup[d.pos-1] = true
        }
        if d.pos+1 <= 500000 && !dup[d.pos+1] {
            q.Push(Data{d.pos + 1, d.cnt + 1})
            dup[d.pos+1] = true
        }
        if d.pos*2 <= 500000 && !dup[d.pos*2] {
            q.Push(Data{d.pos * 2, d.cnt + 1})
            dup[d.pos*2] = true
        }
    }
}
