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
func PopFront(l *list.List) Pair {
    f := l.Front()
    l.Remove(f)
    return f.Value.(Pair)
}

func PopBack(l *list.List) Pair {
    b := l.Back()
    l.Remove(b)
    return b.Value.(Pair)
}

func Peek(l *list.List) Pair {
    b := l.Back()
    return b.Value.(Pair)
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
    dx := []int{0, 0, -1, +1}
    dy := []int{-1, +1, 0, 0}

    was := make([][]bool, n)
    for i := 0; i < n; i++ {
        was[i] = make([]bool, m)
        for j := 0; j < m; j++ {
            was[i][j] = false
        }
    }

    P := func(x, y int) bool {
        return x >= 0 && y >= 0 && x < n && y < m
    }

    cnt, minus := 0, 0

    var dfs func(x, y int)
    dfs = func(x, y int) {

        for k := 0; k < 4; k++ {
            i, j := x + dx[k], y + dy[k]
            if P(i, j) && a[i][j] == a[x][y] {
                minus++
            }
        }

        for k := 0; k < 4; k++ {
            i, j := x + dx[k], y + dy[k]
            if !P(i, j) || was[i][j] || a[i][j] != a[x][y] {
                continue
            }

            was[i][j] = true
            cnt++
            dfs(i, j)
        }
    }

    res := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if was[i][j] {
                continue
            }
            cnt, minus = 1, 0
            was[i][j] = true
            dfs(i, j)
            per := (cnt * 4) - minus
            res += per * cnt
        }
    }

    fprintf("RES = %d\n", res)





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
    dx := []int{0, 0, -1, +1}
    dy := []int{-1, +1, 0, 0}
    was := make([][]bool, n)
    for i := 0; i < n; i++ {
        was[i] = make([]bool, m)
        for j := 0; j < m; j++ {
            was[i][j] = false
        }
    }

    type Corner struct {
        dx []int
        dy []int
    }

    corners := []Corner{
        {
            dx: []int{-1, 0},
            dy: []int{0, -1},
        },
        {
            dx: []int{-1, 0},
            dy: []int{0, +1},
        },
        {
            dx: []int{+1, 0},
            dy: []int{0, -1},
        },
        {
            dx: []int{+1, 0},
            dy: []int{0, +1},
        },
    }

    // internal := []Corner{
    //     {
    //         dx: []int{-1, 0},
    //         dy: []int{0, -1},
    //     },
    //     {
    //         dx: []int{-1, 0},
    //         dy: []int{0, +1},
    //     },
    //     {
    //         dx: []int{+1, 0},
    //         dy: []int{0, -1},
    //     },
    //     {
    //         dx: []int{+1, 0},
    //         dy: []int{0, +1},
    //     },
    // }

    P := func(x, y int) bool {
        return x >= 0 && y >= 0 && x < n && y < m
    }

    cnt := 0
    cor := 0

    var dfs func(x, y int)
    dfs = func(x, y int) {

        cnt++
        was[x][y] = true

        for _, corner := range corners {
            ok := true
            internal := false
            first, second := 0, 0
            for k := 0; k < 2; k++ {
                nx, ny := x + corner.dx[k], y + corner.dy[k]
                if corner.dx[k] != 0 {
                    first = corner.dx[k]
                } 
                if corner.dy[k] != 0 {
                    second = corner.dy[k]
                }

                if P(nx, ny) && a[nx][ny] == a[x][y] {
                    ok = false
                }

            }

            xx, yy := x + first, y + second
            if P(x + first, y) && a[x + first][y] == a[x][y] && 
            P(x, y + second) && a[x][y + second] == a[x][y] && 
            P(xx, yy) && a[xx][yy] != a[x][y] {
                internal = true
            }

            if ok || internal {
                cor++
            } 

        }


        for k := 0; k < 4; k++ {
            nx, ny := x + dx[k], y + dy[k]
            if !P(nx, ny) || was[nx][ny] || a[nx][ny] != a[x][y] {
                continue
            }

            dfs(nx, ny)
        }

    }

    res := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if !was[i][j] {
                cnt, cor = 0, 0
                dfs(i, j)
                res += cnt * cor
            }
        }
    }

    fprintf("RES = %d\n", res)






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




