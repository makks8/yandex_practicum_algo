package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	streetLen := readInt(scanner)
	houseNumbers := readArray(scanner)

	nearestZeroList := getNearestZero(houseNumbers, streetLen)

	printArray(nearestZeroList)
}

func getNearestZero(houseNumbers []int, streetLen int) []int {
	nearestZeroList := make([]int, 0, streetLen)
	var zeroFound, isCurrentHouseNumberZero bool
	var houseCounter int
	for _, currentHouseNumber := range houseNumbers {
		isCurrentHouseNumberZero = currentHouseNumber == 0

		if isCurrentHouseNumberZero && !zeroFound {
			for i := houseCounter; i > 0; i-- {
				nearestZeroList = append(nearestZeroList, i)
			}
			zeroFound = true
			houseCounter = 0
		}

		if isCurrentHouseNumberZero && zeroFound {
			if houseCounter > 0 {
				middle := houseCounter / 2
				for i := 1; i <= middle; i++ {
					nearestZeroList = append(nearestZeroList, i)
				}
				if houseCounter%2 != 0 {
					nearestZeroList = append(nearestZeroList, middle+1)
				}
				for i := middle; i > 0; i-- {
					nearestZeroList = append(nearestZeroList, i)
				}
				houseCounter = 0
			}
			nearestZeroList = append(nearestZeroList, 0)
		}

		if !isCurrentHouseNumberZero {
			houseCounter++
		}
	}
	for i := 1; i <= houseCounter; i++ {
		nearestZeroList = append(nearestZeroList, i)
	}
	return nearestZeroList
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
