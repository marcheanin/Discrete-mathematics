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

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	sorting(0, n, less, swap)
}

func sorting(l, r int, less func(i, j int) bool, swap func(i, j int)) {
	var i = l
	var j = r
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
