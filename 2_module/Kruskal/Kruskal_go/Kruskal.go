package main

import (
	"fmt"
	"math"
	"sort"
)

type edge struct {
	x, y int
	dist float64
}

type point struct {
	first, second int32
}

func count_dist(a, b point) float64 {
	return math.Sqrt(float64((a.first-b.first)*(a.first-b.first) + (a.second-b.second)*(a.second-b.second)))
}

func main() {
	var n int
	fmt.Scan(&n)
	points := make([]point, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&points[i].first, &points[i].second)
	}
	edges := make([]edge, 0, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{x: i, y: j, dist: count_dist(points[i], points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) (less bool) {
		return edges[i].dist < edges[j].dist
	})
	tree_id := make([]int, n, n)
	for i := 0; i < n; i++ {
		tree_id[i] = i
	}
	m := len(edges)
	var cost float64 = 0
	for i := 0; i < m; i++ {
		a := edges[i].x
		b := edges[i].y
		l := edges[i].dist
		if tree_id[a] != tree_id[b] {
			cost += l
			old_id := tree_id[b]
			new_id := tree_id[a]
			for j := 0; j < n; j++ {
				if tree_id[j] == old_id {
					tree_id[j] = new_id
				}
			}
		}
	}
	fmt.Printf("%.2f", cost)
}
