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
	sorting(0, n, less, swap)
}

func sorting(l, r int, less func(i, j int) bool, swap func(i, j int)) {
	if r-l <= 1 {
		return
	}

	var i = l
	var j = r - 1
	//var x = arr[(l+r)/2]

	for i <= j {
		for less(i, (l+r)/2) { // arr[i] < x
			i++
		}
		for less((l+r)/2, j) { //arr[j] > x
			j--
		}
		if i <= j {
			swap(i, j)
			i++
			j--
		}
	}
	if j > l {
		sorting(l, j+1, less, swap)
	}
	if i < r {
		sorting(j+1, r, less, swap)
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
