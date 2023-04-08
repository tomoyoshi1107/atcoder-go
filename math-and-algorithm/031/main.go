package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve031()
}

var N int
var dp []int

var sc = bufio.NewScanner(os.Stdin)

func solve031() {
	buf := make([]byte, 2000000)
	sc.Buffer(buf, 2000000)
	sc.Split(bufio.ScanWords)

	N = scanInt()

	dp = make([]int, N+1)
	dp[0] = 0
	for i := 1; i <= N; i++ {
		a := scanInt()
		if i == 1 {
			dp[i] = a
		} else {
			dp[i] = max(dp[i-2]+a, dp[i-1])
		}

	}

	fmt.Print(dp[N])

}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
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
