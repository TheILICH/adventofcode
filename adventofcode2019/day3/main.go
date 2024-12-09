package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"

	// "io"
	// "sort"
	// "strings"
	// "strconv"
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
	
}

func second() {

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
		first()
	}

}
