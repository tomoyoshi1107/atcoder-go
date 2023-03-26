package main

import (
	"fmt"
	"io"
	"os"
)

var (
	input io.Reader
)

func init() {
	input = os.Stdin
}

func main() {
	solve020()
}

func solve020() {
	var n int
	fmt.Fscan(input, &n)
	//スライスの宣言 make([]型名, capacity)  make([]型名, length, capacity)
	buf := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &buf[i])
	}
	var ans int = 0
	for i:=0;i<=n;i++{
		for j :=i+1;j<n;j++{
			for k:=j+1;k<n;k++{
				for l:=k+1;l<n;l++{
					for m:=l+1;m<n;m++{
						if buf[i]+buf[j]+buf[k]+buf[l]+buf[m] == 1000{
							ans++
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}