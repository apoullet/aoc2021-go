package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
    "strings"
)

type bingoBoards struct {
    actual [][25]int
    backing [][]int
}

type bingoBoard struct {
    actual [25]int
    backing []int
}

func (boards *bingoBoards) AddBoard(board [25]int) {
    (*boards).actual = append(boards.actual, board)

    maxVal := getMaxVal(board)
    backingBoard := make([]int, maxVal+1)

    for i := range(backingBoard) {
        backingBoard[i] = -1
    }

    for _, val := range(board) {
        backingBoard[val] = 0
    }

    (*boards).backing = append(boards.backing, backingBoard)
}

func (boards bingoBoards) IsWinning() (bool, bingoBoard) {
    for index, currentBoard := range boards.actual {
        backingBoard := boards.backing[index]

        // First Column
        if backingBoard[currentBoard[0]] + backingBoard[currentBoard[5]] + backingBoard[currentBoard[10]] + backingBoard[currentBoard[15]] + backingBoard[currentBoard[20]] == 5 {
            return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // Second Column
        if backingBoard[currentBoard[1]] + backingBoard[currentBoard[6]] + backingBoard[currentBoard[11]] + backingBoard[currentBoard[16]] + backingBoard[currentBoard[21]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // Third Column
        if backingBoard[currentBoard[2]] + backingBoard[currentBoard[7]] + backingBoard[currentBoard[12]] + backingBoard[currentBoard[17]] + backingBoard[currentBoard[22]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // Fourth Column
        if backingBoard[currentBoard[3]] + backingBoard[currentBoard[8]] + backingBoard[currentBoard[13]] + backingBoard[currentBoard[18]] + backingBoard[currentBoard[23]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // Fifth Column
        if backingBoard[currentBoard[4]] + backingBoard[currentBoard[9]] + backingBoard[currentBoard[14]] + backingBoard[currentBoard[19]] + backingBoard[currentBoard[24]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // First Row
        if backingBoard[currentBoard[0]] + backingBoard[currentBoard[1]] + backingBoard[currentBoard[2]] + backingBoard[currentBoard[3]] + backingBoard[currentBoard[4]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // Second Row
        if backingBoard[currentBoard[5]] + backingBoard[currentBoard[6]] + backingBoard[currentBoard[7]] + backingBoard[currentBoard[8]] + backingBoard[currentBoard[9]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // Third Row
        if backingBoard[currentBoard[10]] + backingBoard[currentBoard[11]] + backingBoard[currentBoard[12]] + backingBoard[currentBoard[13]] + backingBoard[currentBoard[14]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // Fourth Row
        if backingBoard[currentBoard[15]] + backingBoard[currentBoard[16]] + backingBoard[currentBoard[17]] + backingBoard[currentBoard[18]] + backingBoard[currentBoard[19]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
        // Fifth Row
        if backingBoard[currentBoard[20]] + backingBoard[currentBoard[21]] + backingBoard[currentBoard[22]] + backingBoard[currentBoard[23]] + backingBoard[currentBoard[24]] == 5 {
               return true, bingoBoard{actual: currentBoard, backing: boards.backing[index]}
        }
    }

    return false, bingoBoard{actual: [25]int{}, backing: []int{}}
}

func (boards *bingoBoards) markBoards(value int) {
    for _, board := range boards.backing {
        if len(board) <= value {
            continue
        }

        board[value] = 1
    }
}

func main() {
	fmt.Println(part1())
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
        boards.markBoards(value)

        isWinning, board := boards.IsWinning()

        if isWinning {
            unmarkedNumbersSum := 0

            for _, val := range board.actual {
                if board.backing[val] == 0 {
                    unmarkedNumbersSum += val
                }
            }

            return unmarkedNumbersSum * value
        }
    }

    return -1
}

func loadBoards(scanner *bufio.Scanner) bingoBoards {
    boards := bingoBoards{actual: [][25]int{}, backing: [][]int{}} 

    currentBoard := [25]int{}
    var rowIndex int

	for scanner.Scan() {
        currentLine := scanner.Text()
        
        if currentLine == "" {
            boards.AddBoard(currentBoard)
            currentBoard = [25]int{}
            rowIndex = 0
            continue
        }

        numbers := removeEmptyVals(strings.Split(currentLine, " "))

        for i := 0; i < 5; i++ {
            value, _ := strconv.Atoi(numbers[i])
            currentBoard[i + rowIndex * 5] = value
        }

        rowIndex++
	}

    boards.AddBoard(currentBoard)
    return boards
}

func getMaxVal(board [25]int) int {
    currentMax := -1

    for _, val := range board {
        if currentMax < val {
            currentMax = val
        }
    }

    return currentMax
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
