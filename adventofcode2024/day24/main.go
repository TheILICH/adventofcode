package main

import (
    // "container/heap"
    "container/list"

    // "slices"
    "sort"
    
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

const (
    inf int = 1e18
    mod int = 16777216
)

func mult(a, b int) int {
    return ((a % mod) * (b % mod)) % mod
}

func add(a, b int) int {
    return ((a % mod) + (b % mod)) % mod
}

func sub(a, b int) int {
    return ((a%mod - b%mod) + mod) % mod
}

func bpow(x, k int) int {
    if k == 0 {
        return 1
    }
    if k%2 == 0 {
        r := bpow(x, k/2)
        return mult(r, r)
    }
    r := bpow(x, k-1)
    return mult(x, r)
}

func inv(n int) int {
    return bpow(n, mod-2)
}

func divide(a, b int) int {
    return mult(a, inv(b))
}

func first() {

    a := make(map[string]int)
    var b [][]string
    g := make(map[string]bool)
    first := true
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }
        
        line = strings.TrimSpace(line)
        if len(line) == 0 {
            first = false
            continue
        }

        s := strings.Split(line, " ")
        if first {
            k, v := s[0], s[1]
            k = k[:len(k) - 1]
            val := int(v[0] - '0')
            g[k] = true
            a[k] = val
            // fprintf("k, v := %s, %d\n", k, val)
        } else {
            fr, op, to, res := s[0], s[1], s[2], s[4]
            g[fr] = true
            g[to] = true
            g[res] = true
            b = append(b, []string{fr, op, to, res})
        }
    }

    check := func() bool {
        for k := range g {
            if _, ex := a[k]; !ex {
                return false
            }
        }

        return true
    }

    for !check() {

        for _, s := range b {

            fr, op, to, res := s[0], s[1], s[2], s[3]
            if _, ex := a[res]; ex {
                continue
            }

            x, one := a[fr]
            y, two := a[to]
            
            if !one || !two {
                continue
            }

            if op == "XOR" {
                x ^= y
            } else if op == "AND" {
                x &= y
            } else {
                x |= y
            }

            a[res] = x
        }
    }

    var z []string
    for k := range a {
        if k[0] != 'z' {
            continue
        }

        z = append(z, k)
    }

    sort.Strings(z)
    ans := ""
    for _, zz := range z {
        ans = strconv.Itoa(a[zz]) + ans
    }

    final, _ := strconv.ParseInt(ans, 2, 64)

    fprintf("FINAL = %d and ANS = %s\n", final, ans)


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
 
    // fout, _ := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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


