package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	rdr *bufio.Reader
)

const (
	BUFSIZE = 10000000
	MOD     = 1000000007
)

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	solve022()
}
func solve022() {
	n := readInt()
	buf := readIntSlice()
	cnt := make(map[int]int)
	for i := 0; i <= n; i++ {
		cnt[buf[i]]++
	}
	ans := 0
	for i := 1; i <= 49999; i++ {
		v1, flg1 := cnt[50000+i]
		v2, flg2 := cnt[50000-i]
		if flg1 && flg2 {
			ans += v1 * v2
		}
	}
	ans += (cnt[50000] * (cnt[50000] - 1)) / 2
	fmt.Printf("%d\n", ans)
}

func readInt() int {
	return s2i(readline())
}

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2i(v))
	}
	return slice
}

func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}

func i2s(i int) string {
	return strconv.Itoa(i)
}

// bool <-> int
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func i2b(i int) bool {
	return i != 0
}
