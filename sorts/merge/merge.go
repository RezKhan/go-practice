package main

import (
	"fmt"
	"os"
	rwf "readwritefile"
)

// takes the 2 component arrays, compares and merges them
func merge(leftArray []int64, rightArray []int64) []int64 {
	var sortedArray []int64
	i := 0
	j := 0
	leftLen := len(leftArray)
	rightLen := len(rightArray)

	// sort left and right
	for (i < leftLen) && (j < rightLen) {
		if leftArray[i] < rightArray[j] {
			sortedArray = append(sortedArray, leftArray[i])
			i++
		} else {
			sortedArray = append(sortedArray, rightArray[j])
			j++
		}
	}

	// fill remainder
	for ; i < leftLen; i++ {
		sortedArray = append(sortedArray, leftArray[i])
	}
	for ; j < rightLen; j++ {
		sortedArray = append(sortedArray, rightArray[j])
	}
	return sortedArray
}

// splits the unsorted array into left and right halves
// unlike python, in go we keep the midpoint instead of increasing by one
func mergeSort(unsortedArray []int64) []int64 {
	if len(unsortedArray) <= 1 {
		return unsortedArray
	} else {
		midpoint := len(unsortedArray) / 2

		leftArray := unsortedArray[:midpoint]
		rightArray := unsortedArray[midpoint:]

		left := mergeSort(leftArray)
		right := mergeSort(rightArray)
		return merge(left, right)
	}
}

func main() {
	var submitFile string
	if len(os.Args) != 2 {
		fmt.Println("Specificy filename to sort. ")
		return
	} else {
		submitFile = os.Args[1]
	}
	unsortedArray, readErr := rwf.ReadFile(submitFile)
	if readErr != nil {
		return
	}
	sortedArray := mergeSort(unsortedArray)
	sortedFile := fmt.Sprintf("mergesort-%s", submitFile)
	_, writeError := rwf.WriteFile(sortedFile, sortedArray)
	if writeError != nil {
		println("Could not write")
	}
}
