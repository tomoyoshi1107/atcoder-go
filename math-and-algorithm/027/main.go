package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve027()
}

var N int
var err error
var a, c []int

var sc = bufio.NewScanner(os.Stdin)

func solve027() {
	buf := make([]byte, 2000000)
	sc.Buffer(buf, 2000000)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	N, err = strconv.Atoi(sc.Text())

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	a = make([]int, N)
	c = make([]int, N)

	for i := 0; i < N; i++ {
		a[i] = scanInt()
	}

	mergeSort(0, N)

	s := make([]string, N)
	for i, v := range c {
		s[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(s, " "))

}

func mergeSort(l, r int) {
	//r-l =1の場合、すでにソートされているのでなにもしない
	if r-l == 1 {
		return
	}
	//二つに分割した後、小さい配列をソート
	m := (l + r) / 2
	mergeSort(l, m)
	mergeSort(m, r)

	//この時点で以下の二つの配列がソートされている
	//列Ａ'に相当するもの[A[l],A[l+1],......,A[m-1]]
	//列Bに相当するもの[A[m],A[m+1],......,A[r-1]]
	//以下がmerge操作

	c1, c2, cnt := l, m, 0
	for c1 != m || c2 != r {
		if c1 == m {
			//列A'がからの場合
			c[cnt] = a[c2]
			c2++
		} else if c2 == r {
			//列Bがからの場合
			c[cnt] = a[c1]
			c1++
		} else {
			if a[c1] > a[c2] {
				c[cnt] = a[c2]
				c2++
			} else {
				c[cnt] = a[c1]
				c1++
			}
		}
		cnt++
	}
	//A',Bを合併した配列CをAに移す
	for i := 0; i < cnt; i++ {
		a[l+i] = c[i]
	}

}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	return i
}

func scanFloat64() float64 {
	sc.Scan()
	i, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	return i
}

func parseInt(s string, n int) []int {
	var err error
	tmp := strings.Split(s, " ")
	slice := make([]int, n)
	for i, v := range tmp {
		slice[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error:", err)
			panic(err)
		}
	}
	return slice
}

func parseFloat(s string, n int) []float64 {
	var err error
	tmp := strings.Split(s, " ")
	slice := make([]float64, n)
	for i, v := range tmp {
		slice[i], err = strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("Error:", err)
			panic(err)
		}
	}
	return slice
}
