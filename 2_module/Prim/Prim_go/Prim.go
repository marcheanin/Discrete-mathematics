package main

import "fmt"

type pair struct {
	first, second int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	g := make([][]pair, n, n)
	for i := range g {
		g[i] = make([]pair, 0, 0)
	}
	var x, y, dist int
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y, &dist)
		g[x] = append(g[x], pair{y, dist})
		g[y] = append(g[y], pair{x, dist})
	}
	used := make([]bool, n, n)
	used_v := 1
	ans := 0
	used[0] = true
	for ; used_v < n; used_v++ {
		mun := 10000000
		vmun := 0
		for i := 0; i < n; i++ {
			if used[i] {
				for j := range g[i] {
					if !used[g[i][j].first] {
						if g[i][j].second < mun {
							mun = g[i][j].second
							vmun = g[i][j].first
						}
					}
				}
			}
		}
		used[vmun] = true
		ans += mun
	}
	fmt.Println(ans)
}
