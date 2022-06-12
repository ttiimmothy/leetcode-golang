package main

import "fmt"

func bubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for m := 0; m < len(array)-i-1; m++ {
			if array[m] > array[m+1] {
				array[m], array[m+1] = array[m+1], array[m]
			}
		}
	}
	return array
}

func main() {
	var array = []int{1, 5, 6, 3, 1, 6, 3}
	fmt.Println(bubbleSort(array))
}
