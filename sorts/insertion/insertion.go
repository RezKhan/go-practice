package main

import (
	"fmt"
	"os"
	rwf "readwritefile"
)


func insertionSort(sortingArray []int64) []int64 {
	length := len(sortingArray)
	
	for i := 1; i < length; i++ {
		testValue := sortingArray[i]
		j := i - 1
		for j > 0 && testValue < sortingArray[j] {
			sortingArray[j + 1] = sortingArray[j]
			j -= 1
		}
		sortingArray[j + 1] = testValue
	}
	return sortingArray
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
	sortedArray := insertionSort(unsortedArray)
	sortedFile := fmt.Sprintf("insertionsort-%s", submitFile)

	_, writeError := rwf.WriteFile(sortedFile, sortedArray)
	if writeError != nil {
		println("Could not write")
	}
}
