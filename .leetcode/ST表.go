package main

import (
	"fmt"
	"math"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//输入：
	//8 8
	//9 3 1 7 5 6 0 8
	//1 6
	//1 5
	//2 7
	//2 6
	//1 8
	//4 8
	//3 7
	//1 8
	//
	//输出：
	//9
	//9
	//7
	//7
	//9
	//8
	//7
	//9
	var n, m int
	fmt.Scanln(&n, &m)
	var list = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&list[i])
	}
	var request = make(map[int][]int, m)
	var x, y int
	for i := 0; i < m; i++ {
		fmt.Scanln(&x, &y)
		request[i] = []int{x, y}
	}

	var f [][]int
	f = append(f, make([]int, 17))
	for i := 1; i <= n; i++ {
		tmp := make([]int, 17)
		f = append(f, tmp)
		f[i][0] = list[i]
	}
	for j := 1; j < 17; j++ {
		for i := 1; i+1<<uint(j)-1 <= n; i++ {
			f[i][j] = Max(f[i][j-1], f[i+1<<uint(j-1)][j-1])
		}
	}
	for i := 0; i < m; i++ {
		x, y := request[i][0], request[i][1]
		t := uint(math.Log2(float64(y - x + 1)))
		fmt.Println(Max(f[x][t], f[y+1-1<<t][t]))
	}
}
