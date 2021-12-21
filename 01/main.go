package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    file, _ := os.Open("input.txt")

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    
    parsedLines := []int{}

    for scanner.Scan() {
        value, _ := strconv.Atoi(scanner.Text())
        parsedLines = append(parsedLines, value)
    }

    increasesCount := 0

    for i := 0; i < len(parsedLines)-1; i++ {
        if parsedLines[i] < parsedLines[i+1] { 
            increasesCount++ 
        }
    }

    fmt.Println("Answer 1:", increasesCount)

    increasesCount = 0

    for i := 0; i < len(parsedLines)-3; i++ {
        if parsedLines[i] < parsedLines[i+3] { 
            increasesCount++ 
        }
    }

    fmt.Println("Answer 2:", increasesCount)
}
