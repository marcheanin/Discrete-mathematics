package main

import (
	"fmt"
)

func add(a, b []int32, p int) []int32 {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	var x int32 = 0
	if len(a) != len(b) {
		if len(a) < len(b) {
			a, b = b, a
		}

	}
	k := len(b) - 1
	for i := len(a) - 1; i >= 0; i-- {
		if k >= 0 {
			a[i] += b[k]
		}
		a[i] += x
		x = 0
		if a[i] >= int32(p) {
			x = a[i] / int32(p)
			a[i] %= int32(p)
		}
		k--
	}

	if x != 0 {
		var res []int32
		res = append(res, x)
		for j := 0; j < len(a); j++ {
			res = append(res, a[j])
		}
		a = res
	}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	a := []int32{1}
	b := []int32{1}
	fmt.Print(add(a, b, 2))
}
