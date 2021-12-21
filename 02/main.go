package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    part1()
    part2()
}

func part1() {
    file, _ := os.Open("input.txt")

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    distance := 0
    depth := 0
    
    for scanner.Scan() {
        command := strings.Split(scanner.Text(), " ")
        value, _ := strconv.Atoi(command[1])

        switch command[0] {
        case "forward":
            distance += value
        case "up":
            depth -= value
        case "down":
            depth += value
        }
    }

    fmt.Println("Answer 1:", distance * depth)
}

func part2() {
    file, _ := os.Open("input.txt")

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    distance := 0
    depth := 0
    aim := 0
    
    for scanner.Scan() {
        command := strings.Split(scanner.Text(), " ")
        value, _ := strconv.Atoi(command[1])

        switch command[0] {
        case "forward":
            distance += value
            depth += (aim * value)
        case "up":
            aim -= value
        case "down":
            aim += value
        }
    }

    fmt.Println("Answer 2:", distance * depth)
}
