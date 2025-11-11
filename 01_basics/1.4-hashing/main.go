/*
Hashing is a fundamental technique used to uniquely identify and store data in a way that allows for extremely fast retrieval. The average time complexity for insertion, deletion, and search operations is O(1) (constant time).Its core idea is to use a special function, called a hash function, to convert an input (like a string, number, or object) into a fixed-size integer. This integer is then used as an index to store the data in an array called a hash table.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// countFrequency([]int{1,2,3,3,4})
	array1 := []int{10, 5, 10, 15, 10, 5}
	printMinMaxFreq(array1)
}

func countFrequency(array []int){
	size := len(array)
	visited := make([]bool,size) // make command initialized bool as false automatically
	for i:= 0 ; i < size; i ++{
		if visited[i]{
			continue
		}
		count:= 1
		visited[i] = true

		for j:=i+1; j < size; j++{
			if array[i] == array[j]{
				count++
				visited[j] = true
			}
		}

		fmt.Printf("%d %d \n",array[i],count)
	}
}

func printMinMaxFreq(arr[]int){
	if len(arr) == 0{
		return
	}

	counts := make(map[int]int)
	for _,num := range arr{
		counts[num]++
	}

	minFreq := math.MaxInt32
	maxFreq := 0

	var minFreqNum int
	var maxFreqNum int

	for num, freq := range counts {
		if freq > maxFreq {
			maxFreq = freq
			maxFreqNum = num
		}

		if freq < minFreq {
			minFreq = freq
			minFreqNum = num
		}
	}

	fmt.Printf("%d %d\n",maxFreqNum, minFreqNum)
}