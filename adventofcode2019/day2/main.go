package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "io"
	// "sort"
	// "strings"
	"strconv"
	// "container/list"
)

const (
	inf int = 1e18
	mod int = 998244353
	N   int = 500
	M   int = 500
)

var (

)

func first() {

	var a []int
	end := false
	for !end {
		line, err := r.ReadString(',')
		if err != nil && err.Error() == "EOF" {
			end = true
		}

		n, _ := strconv.Atoi(strings.Trim(line, ","))
		a = append(a, n)
	}

	a[1] = 12
	a[2] = 2

	for i := 0; i < len(a); i += 4 {
		if a[i] == 99 {
			break
		}

		if a[i] == 1 {
			a[a[i + 3]] = a[a[i + 1]] + a[a[i + 2]]
		} else if a[i] == 2 {
			a[a[i + 3]] = a[a[i + 1]] * a[a[i + 2]]
		}
	}

	fprintf("%d\n", a[0])

}

func Go(a []int, noun, verb int) int {

	a[1] = noun
	a[2] = verb
	n := len(a)

	for i := 0; i < n; i += 4 {
		if a[i] == 99 {
			break
		}

		if a[i] == 1 {
			if a[i + 3] < n && a[i + 1] < n && a[i + 2] < n {
				a[a[i + 3]] = a[a[i + 1]] + a[a[i + 2]]
			} else {
				break
			}
		} else if a[i] == 2 {
			if a[i + 3] < n && a[i + 1] < n && a[i + 2] < n {
				a[a[i + 3]] = a[a[i + 1]] * a[a[i + 2]]
			} else {
				break
			}
		}
	}

	return a[0]

}

func second() {

	var a []int
	end := false
	for !end {
		line, err := r.ReadString(',')
		if err != nil && err.Error() == "EOF" {
			end = true
		}

		n, _ := strconv.Atoi(strings.Trim(line, ","))
		a = append(a, n)
	}

	b := make([]int, len(a))
	copy(b, a)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if Go(a, i, j) == 19690720 {
				fprintf("%d\n", i * 100 + j)
				return
			}
			copy(a, b)	
		}
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

	fin, _ := os.Open("f.in")
	defer fin.Close()
	r = bufio.NewReader(fin)

	// fout, _ := os.Open("f.out")
	// defer fout.Close()
	// w = bufio.NewWriter(fout)

	defer w.Flush()

	tt := 1
	// fscan(&tt)

	for i := 0; i < tt; i++ {
		second()
	}

}



