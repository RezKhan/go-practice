package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
    // initialise with the length of what is provided in arguments or 10k
    var argscounter int64
    var defaultnums int64 = 10000
    
    if len(os.Args) == 2 {
        // confirm if the argument is an integer, otherwise fail out
        temp, err := strconv.ParseInt(os.Args[1], 10, 64)
        if err != nil {
            fmt.Println("Integer not entered: ", err)
            return
        } else {
            argscounter = temp 
        }
    } else {
        argscounter = defaultnums
    }

    filename := fmt.Sprintf("%dnums.txt", argscounter)
    fmt.Println(filename)
    
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Error opening file: ", err)
        return
    }
    defer file.Close()

    for i := 0; i < int(argscounter); i++ {
        randomInteger := rand.Intn(int(argscounter)) + 1

        _, err := fmt.Fprintln(file, randomInteger)
        if err != nil {
            fmt.Println("Error opening file: ", err)
            return
        }
    }
}
