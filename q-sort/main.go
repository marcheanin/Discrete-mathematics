package main

import "fmt"

var arr []int

func swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func less(i, j int) bool {
	return arr[i] < arr[j]
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	sorting(0, n, n/2, less, swap)
}

func sorting(l, r, k int, less func(i, j int) bool, swap func(i, j int)) {
	if r-l <= 1 {
		return
	}
	var x = arr[(l+r)/2]
	var i = l
	var j = r - 1
	for i <= j {
		for arr[i] < x {
			i++
		}
		for arr[j] > x {
			j--
		}
		if i <= j {
			swap(i, j)
			i++
			j--
		}
	}
	if j >= k {
		sorting(l, j+1, k, less, swap)
	} else {
		sorting(j+1, r, k, less, swap)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	qsort(n, less, swap)
	fmt.Println(arr)
}
