package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	houseNumbers := readArray(scanner)

	nearestZeroList := getNearestZero(houseNumbers)

	printArray(nearestZeroList)
}

func getNearestZero(houseNumbers []int) []int {
	var firstZeroIndex, lastZeroIndex int
	var firstZeroFound bool
	var counter = 1
	for index, currentHouse := range houseNumbers {
		if currentHouse == 0 {
			if !firstZeroFound {
				firstZeroIndex = index
				firstZeroFound = true
			}
			lastZeroIndex = index
			counter = 1
			continue
		}
		houseNumbers[index] = counter
		counter++
	}

	var middle int
	for index := lastZeroIndex; index >= 0; index-- {
		if index == firstZeroIndex {
			middle = index
			counter = 1
			continue
		}
		if houseNumbers[index] == 0 {
			middle = houseNumbers[index-1] / 2
			counter = 1
			continue
		}
		if counter > middle {
			continue
		}
		houseNumbers[index] = counter
		counter++
	}

	return houseNumbers
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

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}
