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

type Game struct {
    A Pair
    B Pair
    P Pair
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

    var a []Game
    var A Pair
    var B Pair
    var P Pair
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        if len(line) == 0 {
            continue
        }

        buttonRegex := regexp.MustCompile(`^Button\s([A-Z]):\sX\+(\d{2}),\sY\+(\d{2})$`)
        prizeRegex := regexp.MustCompile(`^Prize:\sX=(\d+),\sY=(\d+)$`)

        if buttonRegex.MatchString(line) {
            match := buttonRegex.FindStringSubmatch(line)
            button := match[1]
            x := match[2]
            y := match[3]
            
            x_int, _ := strconv.Atoi(x)
            y_int, _ := strconv.Atoi(y)

            if button == "A" {
                A = Pair{
                    first: x_int,
                    second: y_int,
                }
            }

            if button == "B" {
                B = Pair{
                    first: x_int,
                    second: y_int,
                }
            }
        } else if prizeRegex.MatchString(line) {

            match := prizeRegex.FindStringSubmatch(line)
            x := match[1]
            y := match[2]

            x_int, _ := strconv.Atoi(x)
            y_int, _ := strconv.Atoi(y)

            P = Pair{
                first: x_int,
                second: y_int,
            }

            a = append(a, Game{
                A: A,
                B: B,
                P: P,
            })
        }

    }



    // for _, g := range a {
    //     fprintf("A.x = %d and A.y = %d\nB.x = %d and B.y = %d\nP.x = %d and P.y = %d\n\n\n", g.A.first, g.A.second, g.B.first, g.B.second, g.P.first, g.P.second)
    // }


    sm := 0
    for _, g := range a {

        mn := inf

        ax, ay, bx, by, px, py := g.A.first, g.A.second, g.B.first, g.B.second, g.P.first, g.P.second
        for i := 0; i < 101; i++ {
            for j := 0; j < 101; j++ {

                axm, aym := ax * i, ay * i
                bxm, bym := bx * j, by * j
                smx, smy := axm + bxm, aym + bym
                if smx == px && smy == py {
                    mn = min(mn, i * 3 + j)
                }

            }
        }

        if mn == inf {
            continue
        }

        sm += mn

    }


    fprintf("SM = %d\n", sm)





}

func second() {

    var a []Game
    var A Pair
    var B Pair
    var P Pair
    end := false
    for !end {

        line, err := r.ReadString('\n')
        if err != nil && err.Error() == "EOF" {
            end = true
        }

        line = strings.TrimSpace(line)
        if len(line) == 0 {
            continue
        }

        buttonRegex := regexp.MustCompile(`^Button\s([A-Z]):\sX\+(\d{2}),\sY\+(\d{2})$`)
        prizeRegex := regexp.MustCompile(`^Prize:\sX=(\d+),\sY=(\d+)$`)

        if buttonRegex.MatchString(line) {
            match := buttonRegex.FindStringSubmatch(line)
            button := match[1]
            x := match[2]
            y := match[3]
            
            x_int, _ := strconv.Atoi(x)
            y_int, _ := strconv.Atoi(y)

            if button == "A" {
                A = Pair{
                    first: x_int,
                    second: y_int,
                }
            }

            if button == "B" {
                B = Pair{
                    first: x_int,
                    second: y_int,
                }
            }
        } else if prizeRegex.MatchString(line) {

            match := prizeRegex.FindStringSubmatch(line)
            x := match[1]
            y := match[2]

            x_int, _ := strconv.Atoi(x)
            y_int, _ := strconv.Atoi(y)
            x_int += tril
            y_int += tril

            P = Pair{
                first: x_int,
                second: y_int,
            }

            a = append(a, Game{
                A: A,
                B: B,
                P: P,
            })
        }

    }



    for _, g := range a {
        fprintf("A.x = %d and A.y = %d\nB.x = %d and B.y = %d\nP.x = %d and P.y = %d\n\n\n", g.A.first, g.A.second, g.B.first, g.B.second, g.P.first, g.P.second)
    }


    sm := 0
    for _, g := range a {

        mn := inf * 9

        ax, ay, bx, by, px, py := g.A.first, g.A.second, g.B.first, g.B.second, g.P.first, g.P.second
        for i := 0; i < 101; i++ {
            for j := 0; j < 101; j++ {

                axm, aym := ax * i, ay * i
                bxm, bym := bx * j, by * j
                smx, smy := axm + bxm, aym + bym
                if smx == px && smy == py {
                    mn = min(mn, i * 3 + j)
                }

            }
        }

        if mn == inf {
            continue
        }

        sm += mn

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

    file_name := "example.txt"
    // file_name := "f.in"

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





