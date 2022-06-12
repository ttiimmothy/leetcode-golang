package main

import "fmt"

func mergeSortGo(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var m = len(array) / 2
	var first = array[:m]
	var second = array[m:]
	return mergeSol(
		mergeSortGo(first),
		mergeSortGo(second))
}

func mergeSol(first []int, second []int) []int {
	var size, i, n = len(first) + len(second), 0, 0
	//make(type,len,capacity)
	newArray := make([]int, size)
	for o := 0; o < len(newArray); o++ {
		if i < len(first) && n > len(second)-1 {
			newArray[o] = first[i]
			i++
		} else if i > len(first)-1 && n < len(second) {
			newArray[o] = second[n]
			n++
		} else if first[i] < second[n] {
			newArray[o] = first[i]
			i++
		} else {
			newArray[o] = second[n]
			n++
		}
	}
	return newArray
}

func main() {
	var array = []int{1, 5, 6, 3, 1, 6, 3}
	fmt.Println(mergeSortGo(array))
}
