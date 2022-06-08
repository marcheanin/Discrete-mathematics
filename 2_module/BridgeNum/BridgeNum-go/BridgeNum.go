package main

import "fmt"

var g, num [][]int
var ans, tin, tup []int
var used []bool
var timer = 0

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(v, predok, nomer int) {
	timer++
	used[v] = true
	tin[v] = timer
	tup[v] = timer
	for i := 0; i < len(g[v]); i++ {
		to := g[v][i]
		if used[to] == false {
			dfs(to, v, num[v][i])
			tup[v] = min(tup[v], tup[to])
		} else if nomer != num[v][i] {
			tup[v] = min(tup[v], tin[to])
		}
	}
	if predok != -1 && tup[v] > tin[predok] {
		ans = append(ans, nomer)
	}
}

func main() {
	var n, m, x, y int
	fmt.Scan(&n, &m)
	g = make([][]int, n, n)
	num = make([][]int, n, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0, 0)
		num[i] = make([]int, 0, 0)
	}
	used = make([]bool, n, n)
	tin = make([]int, n, n)
	tup = make([]int, n, n)
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		num[x] = append(num[x], i+1)
		num[y] = append(num[y], i+1)
	}
	for k := 0; k < n; k++ {
		if !used[k] {
			dfs(k, -1, -1)
		}
	}
	fmt.Println(len(ans))
}

/*
8
10
2 3
5 2
5 7
5 3
4 3
5 6
6 0
1 6
7 4
0 1
*/
