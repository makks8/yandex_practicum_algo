package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := makeScanner()
	yLen := readInt(scanner)
	xLen := readInt(scanner)
	matrix := readMatrix(scanner, yLen, xLen)
	yPosition := readInt(scanner)
	xPosition := readInt(scanner)
	neighbours := GetNeighbours(matrix, yLen, xLen, yPosition, xPosition)

	for _, elem := range neighbours {
		fmt.Print(elem)
		fmt.Print(" ")
	}
}

func GetNeighbours(matrix [][]int, yLen, xLen, yPosition, xPosition int) []int {
	upperNeighbourPosition, lowerNeighbourPosition := yPosition-1, yPosition+1
	leftNeighbourPosition, rightNeighbourPosition := xPosition-1, xPosition+1

	neighbours := make([]int, 0, 4)

	if upperNeighbourPosition >= 0 && upperNeighbourPosition < yLen {
		neighbours = append(neighbours, matrix[upperNeighbourPosition][xPosition])
	}
	if lowerNeighbourPosition < yLen && lowerNeighbourPosition >= 0 {
		neighbours = append(neighbours, matrix[lowerNeighbourPosition][xPosition])
	}
	if leftNeighbourPosition >= 0 && leftNeighbourPosition < xLen {
		neighbours = append(neighbours, matrix[yPosition][leftNeighbourPosition])
	}
	if rightNeighbourPosition < xLen && rightNeighbourPosition >= 0 {
		neighbours = append(neighbours, matrix[yPosition][rightNeighbourPosition])
	}

	sort.Slice(neighbours, func(i, j int) bool {
		return neighbours[i] < neighbours[j]
	})

	return neighbours
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readMatrix(scanner *bufio.Scanner, rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = readArray(scanner)
	}
	return matrix
}
