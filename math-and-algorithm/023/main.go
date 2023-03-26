package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve023()
}
func solve023() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*64*1024)
	sc.Buffer(buf, 64*64*1024)
	//sc.Split(bufio.ScanWords)

	sc.Scan()

	n, err := strconv.Atoi(sc.Text())

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var blue int64
	var red int64
	sc.Scan()
	strB := sc.Text()
	tmpB := strings.Split(strB, " ")
	for _, v := range tmpB {
		tmp, err := strconv.Atoi(v)
		blue += int64(tmp)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	sc.Scan()
	strR := sc.Text()
	tmpR := strings.Split(strR, " ")
	for _, v := range tmpR {
		tmp, err := strconv.Atoi(v)
		red += int64(tmp)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	fmt.Println(float64(blue)/float64(n) + float64(red)/float64(n))

}

func scanInt() int {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		fmt.Println("Error:", err)
	}
	return i
}

func scanFloat64() float64 {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	i, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		fmt.Println("Error:", err)
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
			return slice
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
			return slice
		}
	}
	return slice
}
