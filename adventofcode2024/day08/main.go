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
    a := make([][]byte, n)
    for i := 0; i < n; i++ {
        a[i] = make([]byte, m)
    }

    C := func() {
        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                a[i][j] = b[i][j]
            }
        }
    }

    C()

    d := make(map[byte][]Pair)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if a[i][j] == '.' || a[i][j] == '#' {
                continue
            }

            if _, ex := d[a[i][j]]; !ex {
                d[a[i][j]] = make([]Pair, 0)
            }
            d[a[i][j]] = append(d[a[i][j]], Pair{
                first: i,
                second: j,
            })
        }
    }
    
    P := func(x, y int) bool {
        return x >= 0 && y >= 0 && x < n && y < m
    }

    cnt := 0
    for _, v := range d {
        for i := 0; i < len(v); i++ {
            for j := i + 1; j < len(v); j++ {

                row_diff := v[j].first - v[i].first
                col_diff := v[j].second - v[i].second
                up := Pair{
                    first: v[i].first - row_diff,
                    second: v[i].second - col_diff,
                }

                down := Pair{
                    first: v[j].first + row_diff,
                    second: v[j].second + col_diff,
                }

                if P(up.first, up.second) && a[up.first][up.second] != '#' {
                    cnt++
                    a[up.first][up.second] = '#'
                }
                if P(down.first, down.second) && a[down.first][down.second] != '#' {
                    cnt++
                    a[down.first][down.second] = '#'
                }
            }
        }


    }

    for _, s := range a {
        for _, c := range s {
            fprintf("%c", c)
        }
        fprintf("\n")
    }
    fprintf("\n")

    fprintf("CNT = %d\n", cnt)


}

func second() {

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
    d := make(map[byte][]Pair)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if a[i][j] == '.' || a[i][j] == '#' {
                continue
            }
            if _, ex := d[a[i][j]]; !ex {
                d[a[i][j]] = make([]Pair, 0)
            }
            d[a[i][j]] = append(d[a[i][j]], Pair{
                first: i,
                second: j,
            })
        }
    }

    P := func(x, y int) bool {
        return x >= 0 && y >= 0 && x < n && y < m
    }

    cnt := 0
    for _, v := range d {
        cnt += len(v)
        for i := 0; i < len(v); i++ {
            for j := i + 1; j < len(v); j++ {
                row_diff := v[j].first - v[i].first
                col_diff := v[j].second - v[i].second

                // fprintf("v[i] = {%d %d} and v[j] = {%d %d} and row_diff = %d and col_diff = %d\n", v[i].first, v[i].second, v[j].first, v[j].second, row_diff, col_diff)

                x, y := v[i].first - row_diff, v[i].second - col_diff
                for P(x, y) {

                    if a[x][y] == '.' {
                        a[x][y] = '#'
                        cnt++
                    }

                    x, y = x - row_diff, y - col_diff

                }

                x, y = v[j].first + row_diff, v[j].second + col_diff
                for P(x, y) {

                    if a[x][y] == '.' {
                        a[x][y] = '#'
                        cnt++
                    }

                    x, y = x + row_diff, y + col_diff

                }

            }
        }

        // fprintf("\n\n\n")
    }

    for _, s := range a {
        for _, c := range s {
            fprintf("%c", c)
        }
        fprintf("\n")
    }
    fprintf("\n")


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
 
    // fout, _ := os.OpenFile("f.out", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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

