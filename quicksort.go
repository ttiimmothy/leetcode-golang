package main

import "fmt"

func quicksort(array []int, start int, end int) []int {
	if start < end {
		var pivotPosition = partition(array, start, end)
		quicksort(array, start, pivotPosition-1)
		quicksort(array, pivotPosition+1, end)
	}
	return array
}

func partition(array []int, start int, end int) int {
	var pivot = array[end]
	var m = start - 1
	for i := start; i < end; i++ {
		if array[i] < pivot {
			m++
			swap(array, i, m)
		}
	}
	swap(array, m+1, end)
	return m + 1
}

func swap(array []int, i int, m int) {
	var temp = array[i]
	array[i] = array[m]
	array[m] = temp
}

func main() {
	var array = []int{1, 4, 2, 5, 2, 4, 3}
	var array2 = []int{1, 5, 6, 3, 1, 6, 3}
	fmt.Println(quicksort(array, 0, len(array)-1))
	fmt.Println(quicksort(array2, 0, len(array2)-1))
}
