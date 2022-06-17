package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func dfs(v int, count *int, table [][]Pair, visited []bool, num []int, anum []int) {
	visited[v] = true
	num[*count] = v
	anum[v] = *count
	(*count)++
	for i := 0; i < len(table[v]); i++ {
		to := table[v][i].to
		if visited[to] {
			continue
		}
		dfs(to, count, table, visited, num, anum)
	}
}

type Pair struct {
	to int
	s  string
}

func main() {
	var n, m, q0 int
	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscan(stdin, &n, &m, &q0)
	table := make([][]Pair, n)
	for i := 0; i < n; i++ {
		table[i] = make([]Pair, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(stdin, &table[i][j].to)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(stdin, &table[i][j].s)
		}
	}
	num := make([]int, n)
	anum := make([]int, n)
	visited := make([]bool, n)
	count := 0
	dfs(q0, &count, table, visited, num, anum)

	n = count
	bufstdout := bufio.NewWriter(os.Stdout)
	bufstdout.WriteString(strconv.Itoa(n) + "\n")
	bufstdout.WriteString(strconv.Itoa(m) + "\n" + "0" + "\n")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			bufstdout.WriteString(strconv.Itoa(anum[table[num[i]][j].to]) + " ")
		}
		bufstdout.WriteString("\n")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			bufstdout.WriteString(table[num[i]][j].s + " ")
		}
		bufstdout.WriteString("\n")
	}
	bufstdout.Flush()
}
