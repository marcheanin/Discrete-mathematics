package main

import (
	"fmt"
)

func add(a, b []int32, p int) []int32 {
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
		return res
	}
	return a
}

func main() {
	a := []int32{1}
	b := []int32{9, 9, 9}
	fmt.Print(add(a, b, 10))
}
