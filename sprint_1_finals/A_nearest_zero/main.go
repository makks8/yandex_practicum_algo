package main

import (
	"bufio"
	"os"
	"sort"
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

	zeroFound       bool
	isZeroTempSlice bool

	houseNumbers    []int
	tempSlice       []int
	nearestZeroList []int
}

func (n *NearestZero) refreshTempSlice() {
	n.tempSlice = make([]int, 0, n.streetLen-n.currentKey)
	n.tempSliceLen = 0
}

func (n *NearestZero) run() {
	n.nearestZeroList = make([]int, 0, n.streetLen)
	n.count = 1
	for n.currentKey, n.currentHouseNumber = range n.houseNumbers {
		n.houseNumbersFirstNumberNotZero()

		if len(n.tempSlice) > 0 && n.tempSlice[n.tempSliceLen-1] == 0 && n.currentHouseNumber == 0 {
			n.isZeroTempSlice = true
		}

		if n.isZeroTempSlice && n.currentKey < n.streetLen && n.houseNumbers[n.currentKey+1] != 0 {
			n.nearestZeroList = append(n.nearestZeroList, n.tempSlice...)
			n.isZeroTempSlice = false
			n.refreshTempSlice()
		}

		if n.isZeroTempSlice {
			n.tempSlice = append(n.tempSlice, 0)
			continue
		}

		if n.currentHouseNumber == 0 && !n.isZeroTempSlice {
			if !n.zeroFound {
				n.zeroFound = true
			}
			if n.tempSlice == nil {
				n.refreshTempSlice()
			}

			n.count = 1
			n.tempSliceLen = len(n.tempSlice)

			if n.tempSliceLen <= 2 {
				n.nearestZeroList = append(n.nearestZeroList, n.tempSlice...)
				n.refreshTempSlice()
			}

			if n.tempSliceLen > 2 {
				middleIndex := n.tempSliceLen / 2
				remainder := n.tempSliceLen % 2

				tempSliceBegin := n.tempSlice[:middleIndex]
				tempSliceEnd := make([]int, 0, len(tempSliceBegin))
				tempSliceEnd = append(tempSliceEnd, tempSliceBegin...)

				sort.Slice(tempSliceEnd, func(i, j int) bool {
					return tempSliceEnd[i] > tempSliceEnd[j]
				})

				n.nearestZeroList = append(n.nearestZeroList, tempSliceBegin...)
				if remainder != 0 {
					n.nearestZeroList = append(n.nearestZeroList, middleIndex+1)
				}
				n.nearestZeroList = append(n.nearestZeroList, tempSliceEnd...)

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
	if len(n.tempSlice) > 0 {
		n.nearestZeroList = append(n.nearestZeroList, n.tempSlice...)
	}
}

func (n *NearestZero) houseNumbersFirstNumberNotZero() {
	if n.zeroFound {
		return
	}

	if n.currentHouseNumber != 0 {
		n.nearestZeroList = append(n.nearestZeroList, n.count)
		n.count++
	}

	if n.currentHouseNumber == 0 {
		sort.Slice(n.nearestZeroList[:n.currentKey], func(i, j int) bool {
			return n.nearestZeroList[i] > n.nearestZeroList[j]
		})
		return
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
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}
