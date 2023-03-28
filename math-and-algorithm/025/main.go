package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve025()
}

var sc = bufio.NewScanner(os.Stdin)

func solve025() {
	buf := make([]byte, 0, 64*64*1024)
	sc.Buffer(buf, 64*64*1024)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	n, err := strconv.Atoi(sc.Text())

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = scanInt()
	}
	var ans float64
	for i := 0; i < n; i++ {
		ans += float64(a[i])*1/3 + float64(b[i])*2/3
	}
	fmt.Println(ans)
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
