package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Choice int
	Point  int
}

func main() {
	solve024()
}

var sc = bufio.NewScanner(os.Stdin)

func solve024() {
	n := scanInt()
	questions := make([]Question, n)

	for i := 0; i < n; i++ {
		tmpInt := parseInt(scanString(), 2)
		questions[i].Choice = tmpInt[0]
		questions[i].Point = tmpInt[1]
	}
	ans := 0.0
	for i := 0; i < n; i++ {
		ans += 1.0 * float64(questions[i].Point) / float64(questions[i].Choice)
	}
	fmt.Printf("%.12f\n", ans)

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
