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

    var a []int
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        x, _ := strconv.Atoi(line)
        a = append(a, x)

    }

    find := func(n int) int {
        n = n ^ mult(n, 64)
        n = n ^ (n / 32)
        n = n ^ mult(n, 2048)
        n %= mod

        return n
    }

    sm := 0
    for _, x := range a {

        for i := 0; i < 2000; i++ {
            x = find(x)
        }
        sm += x
    }

    fprintf("SM = %d\n", sm)

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
        x, _ := strconv.Atoi(line)
        a = append(a, x)
    }

    
    // mp_cnt := make(map[int]int)
    // for _, x := range a {
    //     mp_cnt[x]++
    // }

    // cnt_repeted := 0
    // for k, v := range mp_cnt {
    //     if v == 1 {
    //         continue
    //     }
    //     fprintf("%d is repeated!\n", k)
    //     cnt_repeted++
    // }

    // fprintf("CNT_REPEATED = %d\n", cnt_repeted)
    // return



    find := func(n int) int {
        n = n ^ mult(n, 64)
        n = n ^ (n / 32)
        n = n ^ mult(n, 2048)
        n %= mod

        return n
    }

    mp := make(map[string]map[int]int)

    for unique, x := range a {

        b := make([]int, 2001)
        b[0] = x
        for i := 0; i < 2000; i++ {
            x = find(x)
            b[i + 1] = x
        }
        
        for i := 0; i < 2001; i++ {
            b[i] %= 10
        }

        for i := 4; i < 2001; i++ {
            onex := b[i - 3] - b[i - 4]
            twox := b[i - 2] - b[i - 3]
            threex := b[i - 1] - b[i - 2]
            fourx := b[i] - b[i - 1]

            key := fmt.Sprintf("%d,%d,%d,%d", onex, twox, threex, fourx)
            if _, ex := mp[key]; !ex {
                mp[key] = make(map[int]int)
            }

            if _, ex := mp[key][unique]; ex {
                continue
            } else {
                mp[key][unique] = 0
            }

            mp[key][unique] += b[i]
        }
    }

    mx := -inf
    for _, v := range mp {
        curr_sm := 0
        for _, vv := range v {
            curr_sm += vv
        }
        mx = max(mx, curr_sm)
    }


    fprintf("MX = %d\n", mx)



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

