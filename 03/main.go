package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	part1()
}

func part1() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	onesCounter := [12]int{}
	lineCount := 0

	for scanner.Scan() {
		binaryNum := scanner.Text()

		for i := 0; i < len(binaryNum); i++ {
			if binaryNum[i] == 48 {
				onesCounter[i]++
			}
		}

		lineCount++
	}

	gammaRate, epsilonRate := 0, 0

	for i, onesCount := range onesCounter {
		if onesCount > lineCount / 2 {
			gammaRate += int(math.Pow(2, float64(11-i)))
		} else {
			epsilonRate += int(math.Pow(2, float64(11-i)))
		}
	}

	fmt.Println("Answer 1:", gammaRate * epsilonRate)
}
