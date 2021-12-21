package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingoBoard struct {
	actual  [25]int
	backing []int
}

func getMax(board [25]int) int {
	currentMax := -1

	for _, val := range board {
		if currentMax < val {
			currentMax = val
		}
	}

	return currentMax
}

func addBoard(boards []bingoBoard, actualBoard [25]int) []bingoBoard {
	max := getMax(actualBoard)
	backingBoard := make([]int, max+1)

	for i := range backingBoard {
		backingBoard[i] = -1
	}

	for _, val := range actualBoard {
		backingBoard[val] = 0
	}

	board := bingoBoard{actual: actualBoard, backing: backingBoard}
	return append(boards, board)
}

func removeEmptyVals(numbers []string) []string {
	sanitised := []string{}

	for _, val := range numbers {
		if val != "" {
			sanitised = append(sanitised, val)
		}
	}

	return sanitised
}

func loadBoards(scanner *bufio.Scanner) []bingoBoard {
	boards := []bingoBoard{}

	currentBoard := [25]int{}
	var rowIndex int

	for scanner.Scan() {
		currentLine := scanner.Text()

		if currentLine == "" {
			boards = addBoard(boards, currentBoard)
			currentBoard = [25]int{}
			rowIndex = 0
			continue
		}

		numbers := removeEmptyVals(strings.Split(currentLine, " "))

		for i := 0; i < 5; i++ {
			value, _ := strconv.Atoi(numbers[i])
			currentBoard[i+rowIndex*5] = value
		}

		rowIndex++
	}

	boards = addBoard(boards, currentBoard)
	return boards
}

func checkColumn(board bingoBoard, startIndex int) bool {
	isSet := func(index int) int {
		return board.backing[board.actual[index]]
	}

	return isSet(startIndex)+isSet(startIndex+5)+isSet(startIndex+10)+isSet(startIndex+15)+isSet(startIndex+20) == 5
}

func checkRow(board bingoBoard, startIndex int) bool {
	isSet := func(index int) int {
		return board.backing[board.actual[index]]
	}

	return isSet(startIndex)+isSet(startIndex+1)+isSet(startIndex+2)+isSet(startIndex+3)+isSet(startIndex+4) == 5
}

func isWinning(board bingoBoard) bool {
	return checkColumn(board, 0) || checkColumn(board, 1) || checkColumn(board, 2) || checkColumn(board, 3) || checkColumn(board, 4) ||
		checkRow(board, 0) || checkRow(board, 5) || checkRow(board, 10) || checkRow(board, 15) || checkRow(board, 20)
}

func markBoards(boards []bingoBoard, value int) {
	for _, board := range boards {
		if len(board.backing) > value {
			board.backing[value] = 1
		}
	}
}

func part1() int {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	drawnNumbers := removeEmptyVals(strings.Split(scanner.Text(), ","))

	scanner.Scan()

	boards := loadBoards(scanner)

	for _, number := range drawnNumbers {
		value, _ := strconv.Atoi(number)
		markBoards(boards, value)

		for _, board := range boards {
			if isWinning(board) {
				unmarkedNumbersSum := 0

				for _, val := range board.actual {
					if board.backing[val] == 0 {
						unmarkedNumbersSum += val
					}
				}

				return unmarkedNumbersSum * value
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(part1())
}
