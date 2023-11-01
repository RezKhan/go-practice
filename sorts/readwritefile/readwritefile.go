package readwritefile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(submitFile string) ([]int64, error) {
	unsortedFile, fileOpenErr := os.Open(submitFile)
if fileOpenErr != nil {
		fmt.Println("Unable to open file. ", fileOpenErr)
		return nil, fileOpenErr
	}
	defer unsortedFile.Close()

	// scan the file with bufio scanner
	fileScanner := bufio.NewScanner(unsortedFile)
	fileScanner.Split(bufio.ScanLines)
	var unsortedArray []int64

	// append scanned lines to array
	for fileScanner.Scan() {
		temp, scanTextErr := strconv.ParseInt(fileScanner.Text(), 10, 32)
		if scanTextErr != nil {
			fmt.Println("Failed to scan file. ", scanTextErr)
			return nil, scanTextErr
		} else {
			unsortedArray = append(unsortedArray, temp)
		}
	}
	return unsortedArray, nil
}

func WriteFile(sortedFile string, writeArray []int64) (int, error) {
	// create a file handle
	writeFile, fileCreateError := os.Create(sortedFile)
	if fileCreateError != nil {
		fmt.Println("Unable to create file: ", fileCreateError)
		return -1, fileCreateError
	}

	// write each array item to a separate line in the file
	length := len(writeArray)
	for i := 0; i < length; i++ {
		_, fileWriteError := fmt.Fprintln(writeFile, writeArray[i])
		if fileWriteError != nil {
			fmt.Println("Cannot write line")
			return -1, fileWriteError
		}
	}
	return 0, nil
}


