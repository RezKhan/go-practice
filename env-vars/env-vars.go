package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	envStrings := os.Environ()

	for index, estring := range envStrings {
		fmt.Println(index, estring)
	}

	pwdString := os.Getenv("PWD")
	fmt.Println("pwd is : ", pwdString)
}
