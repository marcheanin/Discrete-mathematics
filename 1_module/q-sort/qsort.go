package main

import "fmt"

var arr = []int{
	-385, -303, -259, -268, -330, -347, -376, -253, -259, -376,
	-352, -243, -284, -356, -322, -273, -356, -287, -303, -307,
	-348, -389, -235, -298, -406, -275, -408, -315, -368, -397,
	-341, -364, -398, -275, -334, -290, -376, -269,
}

func swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func less(i, j int) bool {
	return arr[i] < arr[j]
}

//с презентации первого модуля

func partition(l, r int, less func(i, j int) bool, swap func(i, j int)) int {
	i := l
	j := l
	for j < r {
		if less(j, r) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, r)
	return i
}

func sortrec(l int, r int, less func(i, j int) bool, swap func(i, j int)) {
	if l < r {
		q := partition(l, r, less, swap)
		sortrec(l, q-1, less, swap)
		sortrec(q+1, r, less, swap)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	sortrec(0, n-1, less, swap)
}

func sorting(l, r int, less func(i, j int) bool, swap func(i, j int)) {
	//fmt.Println(l, r)
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
		sorting(l, j, less, swap)
	}
	if i < r {
		sorting(i, r, less, swap)
	}
}

func main() {
	qsort(len(arr), less, swap)
	fmt.Println(arr)
}
