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

	n := NearestZero{
		houseNumbers: houseNumbers,
		streetLen:    streetLen,
		tempSlice:    make([]int, 0, streetLen/3),
	}
	n.run()

	printArray(n.nearestZeroList)
}

type NearestZero struct {
	streetLen          int
	currentKey         int
	currentHouseNumber int
	count              int
	tempSliceLen       int

	zeroFound bool

	houseNumbers    []int
	tempSlice       []int
	nearestZeroList []int
}

func (n *NearestZero) refreshTempSlice() {
	n.tempSlice = make([]int, 0, (n.streetLen-n.currentKey)/3)
	n.tempSliceLen = 0
}

func (n *NearestZero) run() {
	n.nearestZeroList = make([]int, 0, n.streetLen)
	n.count = 1
	var middleIndex, remainder int
	for n.currentKey, n.currentHouseNumber = range n.houseNumbers {
		if !n.zeroFound {
			if n.currentHouseNumber != 0 {
				n.tempSlice = append(n.tempSlice, n.count)
				n.count++
				continue
			}

			if n.currentHouseNumber == 0 {
				for i := len(n.tempSlice) - 1; i >= 0; i-- {
					n.nearestZeroList = append(n.nearestZeroList, n.tempSlice[i])
				}
				n.refreshTempSlice()
			}
		}

		if n.currentHouseNumber == 0 {
			if !n.zeroFound {
				n.zeroFound = true
			}

			n.count = 1
			n.tempSliceLen = len(n.tempSlice)

			if n.tempSliceLen <= 1 {
				n.nearestZeroList = append(n.nearestZeroList, n.tempSlice...)
				n.refreshTempSlice()
			}

			if n.tempSliceLen > 1 {
				middleIndex = n.tempSliceLen / 2
				remainder = n.tempSliceLen % 2

				n.tempSlice = n.tempSlice[:middleIndex]
				n.nearestZeroList = append(n.nearestZeroList, n.tempSlice...)
				if remainder != 0 {
					n.nearestZeroList = append(n.nearestZeroList, middleIndex+1)
				}
				for i := middleIndex - 1; i >= 0; i-- {
					n.nearestZeroList = append(n.nearestZeroList, n.tempSlice[i])
				}
				n.refreshTempSlice()
			}
			n.nearestZeroList = append(n.nearestZeroList, 0)
		}

		if n.currentHouseNumber != 0 && n.zeroFound {
			n.tempSlice = append(n.tempSlice, n.count)
			n.tempSliceLen++
			n.count++
		}

	}
	if n.tempSliceLen > 0 {
		n.nearestZeroList = append(n.nearestZeroList, n.tempSlice...)
	}
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
