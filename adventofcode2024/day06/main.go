package main

import (
    // "container/heap"
    "container/list"

    // "slices"
    // "sort"
    
    "os"
    "fmt"
    "bufio"

    "strings"
    // "strconv"
    // "regexp"
    // "math"
    // "cmp"
)

type Item struct {
    val int
}

type Pair struct {
    first int
    second int
}

// deque
func PopFront(l *list.List) int {
    f := l.Front()
    l.Remove(f)
    return f.Value.(int)
}

func PopBack(l *list.List) int {
    b := l.Back()
    l.Remove(b)
    return b.Value.(int)
}

func Peek(l *list.List) int {
    b := l.Back()
    return b.Value.(int)
}

// priority queue
type Heap []Item

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool {

    return h[j].val < h[i].val
}

func (h Heap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(Item))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[0 : n-1]
    return item
}

type PairSlice []Pair

func (p PairSlice) Len() int {
    return len(p) 
}

func (p PairSlice) Less(i, j int) bool {
    if p[i].first == p[j].first {
        return p[i].second < p[j].second
    }
    return p[i].first < p[j].first
}

func (p PairSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func max(a ...int) int {
    res := a[0]
    for _, x := range a {
        if x > res {
            res = x
        }
    }
    return res
}

func min(a ...int) int {
    res := a[0]
    for _, x := range a {
        if x < res {
            res = x
        }
    }
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}

var (
    inf int = 1e9
    mod int = 1e9 + 7
)

func first() {

    var a [][]byte
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)

        a = append(a, []byte(line))
    }

    n, m := len(a), len(a[0])
    sx, sy := -1, -1
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if a[i][j] == '^' {
                sx, sy = i, j
                break
            }
        }
    }

    P := func(x, y int) bool {
        return x >= 0 && y >= 0 && x < n && y < m
    }


    dx := []int{-1, 0, +1, 0}
    dy := []int{0, +1, 0, -1}
    cnt, k := 0, 0

    for P(sx, sy) {

        a[sx][sy] = 'X'
        i, j := sx + dx[k], sy + dy[k]
        if !P(i, j) {
            break
        }
        
        if a[i][j] == '#' {
            k = (k + 1) % 4
            continue
        }

        sx, sy = i, j
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if a[i][j] == 'X' {
                cnt++
            }
        }
    }



    fprintf("CNT = %d\n", cnt)



}

func second() {

    var b [][]byte
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)

        b = append(b, []byte(line))
    }

    n, m := len(b), len(b[0])
    x, y := -1, -1
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if b[i][j] == '^' {
                x, y = i, j
                break
            }
        }
    }

    P := func(x, y int) bool {
        return x >= 0 && y >= 0 && x < n && y < m
    }

    a := make([][]byte, n)
    for i := 0; i < n; i++ {
        a[i] = make([]byte, m)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            a[i][j] = b[i][j]
        }
    }

    dx := []int{-1, 0, +1, 0}
    dy := []int{0, +1, 0, -1}
    cnt := 0
    mx := n * m + 500

    for ii := 0; ii < n; ii++ {
        for jj := 0; jj < m; jj++ {
            a[ii][jj] = '#'

            it := 0
            sx, sy := x, y
            k := 0
            for P(sx, sy) {

                it++
                i, j := sx + dx[k], sy + dy[k]
                if !P(i, j) {
                    break
                }

                if it > mx {
                    cnt++
                    break
                }
                
                if a[i][j] == '#' {
                    k = (k + 1) % 4
                    continue
                }

                sx, sy = i, j
            }

            for i := 0; i < n; i++ {
                for j := 0; j < m; j++ {
                    a[i][j] = b[i][j]
                }
            }


        }
    }



    fprintf("CNT = %d\n", cnt)



}


var (
    w *bufio.Writer
    r *bufio.Reader
)
 
func fscan(a ...any) {
    fmt.Fscan(r, a...)
}
 
func fscanf(format string, a ...any) {
    fmt.Fscanf(r, format, a...)
}
 
func fprintf(format string, a ...any) {
    fmt.Fprintf(w, format, a...)
}
 
func main() {
 
    // r = bufio.NewReader(os.Stdin)
    w = bufio.NewWriter(os.Stdout)

    // file_name := "example.txt"
    file_name := "f.in"
 
    fin, _ := os.Open(file_name)
    defer fin.Close()
    r = bufio.NewReader(fin)
 
    // fout, _ := os.Open("f.out")
    // defer fout.Close()
    // w = bufio.NewWriter(fout)
 
    defer w.Flush()
 
    tt := 1
    // fscan(&tt)
    // fscanf("%d\n", &tt)
 
    for i := 0; i < tt; i++ {
        // first()
        second()
    }
 
}
