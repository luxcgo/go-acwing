package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = int(1e5 + 10)

var (
	n, m     int
	h, e, ne [N]int
	idx      int
	q, d     [N]int
)

func main() {
	s := NewScanner()
	n, m = s.ReadInt(), s.ReadInt()
	for i := 0; i < len(h); i++ {
		h[i] = -1
	}
	for i := 0; i < m; i++ {
		a, b := s.ReadInt(), s.ReadInt()
		add(a, b)
		d[b]++
	}
	if topSort() {
		for i := 0; i < n; i++ {
			fmt.Print(q[i], " ")
		}
	} else {
		fmt.Println(-1)
	}
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func topSort() bool {
	hh := 0
	tt := -1
	for i := 1; i <= n; i++ {
		if d[i] == 0 {
			tt++
			q[tt] = i
		}
	}
	for hh <= tt {
		t := q[hh]
		hh++
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			d[j]--
			if d[j] == 0 {
				tt++
				q[tt] = j
			}
		}
	}
	return tt == n-1
}

/*
 *        __         __
 *       / /_  ___  / /___  ___  __________
 *      / __ \/ _ \/ / __ \/ _ \/ ___/ ___/
 *     / / / /  __/ / /_/ /  __/ /  (__  )
 *    /_/ /_/\___/_/ .___/\___/_/  /____/
 *                /_/
 */

func Unique(a []int) []int {
	j := 0
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] != a[i-1] {
			a[j] = a[i]
			j++
		}
	}
	return a[:j]
}

// Swap 交换两个int指针指向的值
func Swap(a, b *int) {
	*a, *b = *b, *a
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func In() *os.File {
	f, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	return f
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := &Scanner{
		bufio.NewScanner(os.Stdin),
	}
	s.Split(bufio.ScanWords)
	s.Buffer(make([]byte, 1024), int(1e9))
	return s
}

func (s *Scanner) readString() string {
	s.Scan()
	return s.Text()
}

func (s *Scanner) readInt() int {
	i, err := strconv.Atoi(s.readString())
	if err != nil {
		panic(err)
	}
	return i
}

func (s *Scanner) ReadString() string {
	return s.readString()
}

func (s *Scanner) ReadInt() int {
	return s.readInt()
}

func (s *Scanner) ReadStrings(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = s.readString()
	}
	return a
}

func (s *Scanner) ReadInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = s.readInt()
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Memset(a []int, v int) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}
