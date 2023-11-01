package main

import (
	"fmt"
	"os"
	rwf "readwritefile"
)


func bubbleSort(sortingArray []int64) []int64 {
	length := len(sortingArray)

	for length > 1 {
		maxIdx := 0
		for i := 0; i < length-1; i++ {
			if sortingArray[i] > sortingArray[i+1] {
				temp := sortingArray[i]
				sortingArray[i] = sortingArray[i+1]
				sortingArray[i+1] = temp
				maxIdx = i
			}
		}
		length = maxIdx + 1
		// fmt.Println(length)
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
	sortedArray := bubbleSort(unsortedArray)
	sortedFile := fmt.Sprintf("bubblesort-%s", submitFile)

	_, writeError := rwf.WriteFile(sortedFile, sortedArray)
	if writeError != nil {
		println("Could not write")
	}
}
