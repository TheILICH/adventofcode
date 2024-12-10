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

    var a []int
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        n := len(line)
        sm := 0
        for i := 0; i < n; i++ {
            sm += int(line[i] - '0')
        }

        id := 0
        idx := 0
        a = make([]int, sm)
        for i := 0; i < n; i++ {
            x := int(line[i] - '0')
            if i % 2 == 0 {
                for j := 0; j < x; j++ {
                    a[idx] = id
                    idx++
                }
                id++
            } else {
                for j := 0; j < x; j++ {
                    a[idx] = -1
                    idx++
                }
            }
        }

    }

    m := len(a)
    dot := func() int {
        first_dot := 0
        for first_dot < m && a[first_dot] != -1 {
            first_dot++
        }

        if first_dot == m - 1 && a[first_dot] != -1 {
            return -1
        }
        return first_dot
    }

    check := func() bool {

        last_number := m - 1
        for last_number > -1 && a[last_number] == -1 {
            last_number--
        }

        first_dot := dot()

        return first_dot - 1 == last_number

    }

    for i := m - 1; i > -1 && !check(); i-- {

        if a[i] > -1 {
            first_dot := dot()
            a[first_dot], a[i] = a[i], a[first_dot]
        } 

    }

    // fprintf("\na = \n")
    // for i := 0; i < m; i++ {
    //     if a[i] > -1 {
    //         fprintf("%d", a[i])
    //     } else {
    //         fprintf(".")
    //     }
    // }
    // fprintf("\n\n")

    res := 0
    for i := 0; i < m && a[i] > -1; i++ {
        res += a[i] * i
    }

    fprintf("RES = %d\n", res)



}

func second() {

    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)

    }



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

