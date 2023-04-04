package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve028()
}

var N int
var err error
var a, dp []int

var sc = bufio.NewScanner(os.Stdin)

func solve028() {
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
	dp = make([]int, N)

	for i := 0; i < N; i++ {
		a[i] = scanInt()
	}
	for i := 0; i < N; i++ {
		if i == 0 {
			dp[i] = 0
		}
		if i == 1 {
			dp[i] = int(math.Abs(float64(a[i-1] - a[i])))
		}
		if i > 1 {
			v1 := dp[i-1] + int(math.Abs(float64(a[i-1]-a[i])))
			v2 := dp[i-2] + int(math.Abs(float64(a[i-2]-a[i])))
			dp[i] = int(math.Min(float64(v1), float64(v2)))
		}
	}
	fmt.Println(dp[N-1])

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
