package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve009()
}

var N, S int
var dp [][]int
var sc = bufio.NewScanner(os.Stdin)

func solve009() {
	buf := make([]byte, 2000000)
	sc.Buffer(buf, 2000000)
	sc.Split(bufio.ScanWords)

	N = scanInt()

	S = scanInt()
	dp = make([][]int, N+1)
	dp[0] = make([]int, S+1)

	for i := 1; i <= N; i++ {
		dp[i] = make([]int, S+1)
		a := scanInt()
		for j := 0; j <= S; j++ {
			if j < a {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-a]+a)
			}
		}

	}
	ans := "No"
	for i := 0; i <= S; i++ {
		if dp[N][i] == S {
			ans = "Yes"
			break
		}
	}
	fmt.Print(ans)

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
