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
	solve030()
}

var N, W int
var dp [][]int

var sc = bufio.NewScanner(os.Stdin)

func solve030() {
	buf := make([]byte, 2000000)
	sc.Buffer(buf, 2000000)
	sc.Split(bufio.ScanWords)

	N = scanInt()

	W = scanInt()
	dp = make([][]int, N+1)
	//品物iの重さについて上限Wまでスライスを作る
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	dp[0][0] = 0
	var w, v int
	for i := 1; i <= N; i++ {
		w = scanInt()
		v = scanInt()
		for j := 1; j <= W; j++ {

			if j < w { //j< wのとき、品物iは選べない
				dp[i][j] = dp[i-1][j]

			} else { //j >= wのとき、品物iを選べるので選ぶ場合と選ばない場合で価値がおおきくなるほうをとる
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-w]+v)))
			}
		}
	}
	ans := 0
	for i := 0; i <= W; i++ {
		ans = int(math.Max(float64(ans), float64(dp[N][i])))
	}
	fmt.Print(ans)

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
