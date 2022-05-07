package main

import (
	"fmt"
	"sort"
)

var g [][]int

type pair struct {
	first, second int
}

func dfs(v int, used []int, c int) {
	used[v] = c
	for i := range g[v] {
		if used[g[v][i]] == 0 {
			dfs(g[v][i], used, c)
		}
	}
}

type component struct {
	minv int
	v    map[int]bool
	e    []pair
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	g = make([][]int, n, n)
	for i := range g {
		g[i] = make([]int, 0, 0)
	}
	edges := make([]pair, m, m)
	var x, y int
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		edges[i] = pair{x, y}
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	used := make([]int, n, n)
	c := 1
	for i := 0; i < n; i++ {
		if used[i] == 0 {
			dfs(i, used, c)
			c++
		}
	}
	for i := 0; i < n; i++ {
		used[i]--
	}
	s := make([]component, c, c)
	for i := range s {
		s[i].v = make(map[int]bool)
	}
	for i := 0; i < n; i++ {
		(s[used[i]].v)[i] = true
		if i < s[used[i]].minv || len(s[used[i]].v) == 1 {
			s[used[i]].minv = i
		}
	}
	for i := range s {
		for j := range edges {
			_, ok1 := s[i].v[edges[j].first]
			_, ok2 := s[i].v[edges[j].second]
			if ok1 || ok2 {
				s[i].e = append(s[i].e, edges[j])
			}
		}
	}
	sort.Slice(s, func(i, j int) (less bool) {
		if len(s[i].v) != len(s[j].v) {
			return len(s[i].v) > len(s[j].v)
		} else if len(s[i].e) != len(s[j].e) {
			return len(s[i].e) > len(s[j].e)
		} else {
			return s[i].minv < s[j].minv
		}
	})
	ans := s[0]
	for key, _ := range ans.v {
		fmt.Println(key)
	}
	for i := range ans.e {
		fmt.Printf("%d--%d\n", ans.e[i].first, ans.e[i].second)
	}
}
