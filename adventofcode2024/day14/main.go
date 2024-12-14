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
    "strconv"
    "regexp"
    // "math"
    // "cmp"
)

type Robot struct {
    x int
    y int
    vx int
    vy int
}

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
    inf int = 1e18
    mod int = 1e9 + 7
    tril int = 1e13
)

func first() {

    // n, m := 7, 11
    n, m := 103, 101
    var a []Robot
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        re := regexp.MustCompile(`^p=(-?\d+),(-?\d+)\sv=(-?\d+),(-?\d+)$`)
        if match := re.FindStringSubmatch(line); match != nil {

            x, _ := strconv.Atoi(match[2])
            y, _ := strconv.Atoi(match[1])
            vx, _ := strconv.Atoi(match[4])
            vy, _ := strconv.Atoi(match[3])

            a = append(a, Robot{
                x: x,
                y: y,
                vx: vx,
                vy: vy,
            })

        } 
    }

    for i := range a {
        a[i].x += a[i].vx * 100
        a[i].y += a[i].vy * 100

        a[i].x = (a[i].x + (100 * n)) % n
        a[i].y = (a[i].y + (100 * m)) % m
    }

    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, m)
        for j := 0; j < m; j++ {
            res[i][j] = 0
        }
    }

    for i := range a {
        res[a[i].x][a[i].y]++
    }

    one, two, three, four := 0, 0, 0, 0
    for i := 0; i < n / 2; i++ {
        for j := 0; j < m / 2; j++ {
            one += res[i][j]
        }
    }

    for i := 0; i < n / 2; i++ {
        for j := m / 2 + 1; j < m; j++ {
            two += res[i][j]
        }
    }

    for i := n / 2 + 1; i < n; i++ {
        for j := 0; j < m / 2; j++ {
            three += res[i][j]
        }
    }

    for i := n / 2 + 1; i < n; i++ {
        for j := m / 2 + 1; j < m; j++ {
            four += res[i][j]
        }
    }

    fprintf("RES = %d\n", one * two * three * four)




}

func second() {

    // n, m := 7, 11
    n, m := 103, 101
    var a []Robot
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        re := regexp.MustCompile(`^p=(-?\d+),(-?\d+)\sv=(-?\d+),(-?\d+)$`)
        if match := re.FindStringSubmatch(line); match != nil {

            x, _ := strconv.Atoi(match[2])
            y, _ := strconv.Atoi(match[1])
            vx, _ := strconv.Atoi(match[4])
            vy, _ := strconv.Atoi(match[3])

            a = append(a, Robot{
                x: x,
                y: y,
                vx: vx,
                vy: vy,
            })

        } 
    }

    dx := []int{0, 0, -1, +1}
    dy := []int{-1, +1, 0, 0}
    was := make([][]bool, n)
    cnt := 0
    
    C := func() {
        for i := 0; i < n; i++ {
            was[i] = make([]bool, m)
            for j := 0; j < m; j++ {
                was[i][j] = false
            }
        }
    }

    P := func(x, y int) bool {
        return x >= 0 && y >= 0 && x < n && y < m
    }

    var dfs func(x, y int, b [][]byte)
    dfs = func(x, y int, b [][]byte) {

        cnt++
        was[x][y] = true

        for k := 0; k < 4; k++ {
            nx, ny := x + dx[k], y + dy[k]
            if !P(nx, ny) || was[nx][ny] || b[nx][ny] != b[x][y] {
                continue
            }
            dfs(nx, ny, b)
        }

    }

    ans := -1
    for mul := 0; mul < 500; mul++ {
        for i := range a {
            x, y, vx, vy := a[i].x, a[i].y, a[i].vx, a[i].vy

            x += vx * mul
            y += vy * mul

            x = (x + (n * mul)) % n
            y = (y + (m * mul)) % m

            a[i].x = x
            a[i].y = y

        }

        b := make([][]byte, n)
        for i := 0; i < n; i++ {
            b[i] = make([]byte, m)
            for j := 0; j < m; j++ {
                b[i][j] = '.'
            }
        }

        for _, robot := range a {
            b[robot.x][robot.y] = '#'
        }

        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                C()
                cnt = 0
                if !was[i][j] && b[i][j] != '.' {
                    dfs(i, j, b)
                    if cnt > 20 && ans == -1 {
                        ans = mul
                        goto end
                    }
                }
            }
        }

        // fprintf("\n\ncnt = %d\n", cnt)
        // for i := 0; i < n; i++ {
        //     fprintf("%s\n", string(b[i]))
        // }
        // fprintf("\n\n\n\n\n\n\n\n\n\n\n\n\n")
    }


    // fprintf("\n\n\n\nAFTER\n")
    // for _, robot := range a {
    //     fprintf("x, y = %d, %d and vx, vy = %d, %d\n", robot.x, robot.y, robot.vx, robot.vy)
    // }
    // fprintf("\nAFTER\n\n\n\n\n\n\n")

end:
    fprintf("ANS = %d\n", ans)


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
        first()
        // second()
    }
 
}

