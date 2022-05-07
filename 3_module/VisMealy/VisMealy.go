package main

import "fmt"

type state struct {
	in  int
	out string
}

func main() {
	var n, m, q_start int
	fmt.Scanf("%d\n%d\n%d\n", &n, &m, &q_start)
	M := make([][]state, n)
	for i := 0; i < n; i++ {
		M[i] = make([]state, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &M[i][j].in)
		}
		fmt.Scanf("\n")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%s", &M[i][j].out)
		}
		fmt.Scanf("\n")
	}
	fmt.Printf("digraph {\n\trankdir = LR\n")
	//	for i := 0; i < n; i++ {
	//		fmt.Printf("\t%d [shape = circle]\n", i)
	//	}
	//	fmt.Printf("\tdummy -> %d\n", q_start)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("\t%d -> %d [label = \"%c(%s)\"]\n", i, M[i][j].in, 'a'+j, M[i][j].out)
		}
	}
	fmt.Printf("}")
}
