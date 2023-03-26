package main

import (
	"fmt"
)

func main() {
	solve021()
}

func solve021() {
	var n, r int
	fmt.Scan(&n, &r)
	N := 1
	R := 1
	for i := n; i > n-r; i-- {
		N = N * i
	}
	for i := 1; i <= r; i++ {
		R = R * i
	}

	fmt.Print(N / R)

}
