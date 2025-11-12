package main

import "fmt"

func main() {
	fmt.Println("Sorting Algorithms Part 1")
	// fmt.Println(selectionSort([]int{3,2,1,34,33,443,93}))
	fmt.Println(insertionSort([]int{3, 2, 1, 34, 33, 443, 93}))
}

func selectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			temp := array[minIndex]
			array[minIndex] = array[i]
			array[i] = temp
		}
	}
	return array
}

func bubbleSort(array []int) []int {
	n := len(array)
	swapped := false
	for i := n - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return array
}

func insertionSort(array []int) []int {
	for i := range array {
		j := i
		for j > 0 && array[j-1] > array[j] {
			temp := array[j-1]
			array[j-1] = array[j]
			array[j] = temp
			j--
		}
	}
	return array
}
