package main

import "fmt"

func merge(array []int, start int, middle int, end int) {
	//if we need to get the subarray, it will consider the final version of array rather than initial version of array in parameter
	//so this method causes error
	var firstArray = array[start : middle+1]
	var secondArray = array[middle+1 : end+1]
	var length1 = len(firstArray)
	var length2 = len(secondArray)
	var i, j = 0, 0
	var m = start
	for i < length1 && j < length2 {
		if firstArray[i] < secondArray[j] {
			array[m] = firstArray[i]
			i++
		} else {
			array[m] = secondArray[j]
			j++
		}
		m++
	}

	fmt.Println("first ", firstArray)
	fmt.Println("second ", secondArray)
	for i < length1 {
		array[m] = firstArray[i]
		i++
		m++
	}
	for j < length2 {
		array[m] = secondArray[j]
		j++
		m++
	}
	fmt.Println("merge ", array)
}

func mergeSort(array []int, start int, end int) []int {
	if start < end {
		var m = start + (end-start)/2
		//fmt.Println(m)
		fmt.Println("middle index:", m, "middle", array[m])
		mergeSort(array, start, m)
		mergeSort(array, m+1, end)
		merge(array, start, m, end)
	}
	return array
}

func main() {
	//var array = []int{1, 6, 3, 2, 8, 3}
	//var array2 = []int{1, 2, 4, 2, 13, 5, 1, 2, 5, 2}
	var array3 = []int{1, 5, 6, 3, 1, 6, 3}
	//var array4 = []int{12, 11, 13, 5, 6, 7}
	//fmt.Println(mergeSort(array, 0, len(array)-1))
	//fmt.Println(mergeSort(array2, 0, len(array2)-1))
	fmt.Println(mergeSort(array3, 0, len(array3)-1))
	//fmt.Println(mergeSort(array4, 0, len(array4)-1))
}
