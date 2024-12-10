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

    sm := 0
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        s := strings.Split(line, " ")
        n := len(s)

        check, _ := strconv.Atoi(s[0][:len(s[0]) - 1])

        a := make([]int, n - 1)
        for i := 1; i < n; i++ {
            a[i - 1], _ = strconv.Atoi(s[i])
        }
        b := make([]int, n - 1)
        m := n - 1

        C := func() {
            for i := 0; i < m; i++ {
                b[i] = a[i]
            }
        }
        
        stop := false
        for i := 0; i <= (1 << m) && !stop; i++ {
            C()
            for j := 1; j < m; j++ {
                if (1 << (j - 1)) & i != 0 {
                    b[j] *= b[j - 1]
                } else {
                    b[j] += b[j - 1]
                }
            }

            if b[m - 1] == check {
                sm += check
                stop = true
                break
            }
        }


    }

    fprintf("SM = %d\n", sm)


}

func second() {

    sm := 0
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)

        s := strings.Split(line, " ")
        check, _ := strconv.Atoi(s[0][:len(s[0]) - 1])
        
        n := len(s)
        b := make([]int, n - 1)
        for i := 1; i < n; i++ {
            b[i - 1], _ = strconv.Atoi(s[i])
        }
        m := n - 1

        var R func(idx int) bool 
        R = func(idx int) bool {

            if idx == m {
                if b[m - 1] == check {
                    // fprintf("check = %d\n", check)
                    sm += check
                    return true
                }
                
                return false
            }

            og := b[idx]
            
            b[idx] = og + b[idx - 1]
            if R(idx + 1) {
                return true
            }
            
            b[idx] = og * b[idx - 1]
            if R(idx + 1) {
                return true
            }

            b[idx], _ = strconv.Atoi(strconv.Itoa(b[idx - 1]) + strconv.Itoa(og))
            if R(idx + 1) {
                return true
            }
            
            b[idx] = og
            
            return false
            
        }

        R(1)


    }

    fprintf("SM = %d\n", sm)



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
