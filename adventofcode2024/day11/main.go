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

type ListNode struct {
    Val int
    Next *ListNode
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
    inf int = 1e9
    mod int = 1e9 + 7
)

func first() {

    head := &ListNode{
        Val: -1,
    }
    build := head
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        s := strings.Split(line, " ")
        for _, x := range s {
            n, _ := strconv.Atoi(x)
            build.Next = &ListNode{
                Val: n,
            }
            build = build.Next
        }

    }

    for i := 0; i < 25; i++ {
        ch := head.Next
        for ch != nil {
            s := strconv.Itoa(ch.Val)
            ln := len(s)
            if ch.Val == 0 {
                ch.Val = 1
            } else if ln % 2 == 0 {
                left, right := s[0:ln / 2], s[ln / 2: ln]
                left_int, _ := strconv.Atoi(left)
                right_int, _ := strconv.Atoi(right)

                real_next := ch.Next
                ch.Val = left_int
                ch.Next = &ListNode{
                    Val: right_int,
                    Next: real_next,
                }
                ch = ch.Next
            } else {
                ch.Val *= 2024
            }

            ch = ch.Next
        }

    }

    // tr := head.Next
    // for tr != nil {
    //     fprintf("%d ", tr.Val)
    //     tr = tr.Next
    // }
    // fprintf("\n")

    cnt := 0
    tr := head.Next
    for tr != nil {
        cnt++ 
        tr = tr.Next
    }

    fprintf("CNT = %d\n", cnt)





}

func second() {

    var a []int
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        s := strings.Split(line, " ")
        sn := len(s)
        a = make([]int, sn)
        for i := 0; i < sn; i++ {
            a[i], _ = strconv.Atoi(s[i])
        }

    }

    T := func(x int) (int, int) {

        s := strconv.Itoa(x)
        ln := len(s)

        if ln % 2 == 1 {
            return -1, -1
        }

        l, r := s[0:ln / 2], s[ln/2:ln]
        l_int, _ := strconv.Atoi(l)
        r_int, _ := strconv.Atoi(r)

        return l_int, r_int

    }

    M := func(x, y int) string {
        return strconv.Itoa(x) + "," + strconv.Itoa(y)
    }
    
    memo := make(map[string]int)
    var dfs func(x, cnt int) int
    dfs = func(x, cnt int) int {

        if cnt == 0 {
            return 1
        }

        key := M(x, cnt)

        if v, ex := memo[key]; ex {
            return v
        } 

        if x == 0 {
            memo[key] = dfs(1, cnt - 1)
            return memo[key]
        }

        l, r := T(x)
        if l != -1 && r != -1 {
            memo[key] = dfs(l, cnt - 1) + dfs(r, cnt - 1)
            return memo[key]
        }

        memo[key] = dfs(x * 2024, cnt - 1)
        return memo[key]

    }

    sm := 0
    for _, x := range a {
        sm += dfs(x, 75)
    }

    fprintf("SM = %d\n", sm)
    // k := "125,25"
    // fprintf("memo[%s] = %d\n", k, memo[k])





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



