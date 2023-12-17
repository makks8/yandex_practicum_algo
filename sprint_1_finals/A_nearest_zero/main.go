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
	}
	n.run()

	printArray(n.nearestZeroList)
}

type NearestZero struct {
	streetLen          int
	currentKey         int
	currentHouseNumber int
	count              int

	zeroFound bool

	houseNumbers    []int
	nearestZeroList []int
}

func (n *NearestZero) run() {
	n.nearestZeroList = make([]int, 0, n.streetLen)
	var middle int
	for n.currentKey, n.currentHouseNumber = range n.houseNumbers {
		if !n.zeroFound {
			if n.currentHouseNumber != 0 {
				n.count++
				continue
			}

			if n.currentHouseNumber == 0 {
				for i := n.count; i > 0; i-- {
					n.nearestZeroList = append(n.nearestZeroList, i)
				}
				n.zeroFound = true
				n.count = 0
			}
		}

		if n.currentHouseNumber == 0 {
			if n.count > 0 {
				middle = n.count / 2
				for i := 1; i <= middle; i++ {
					n.nearestZeroList = append(n.nearestZeroList, i)
				}
				if n.count%2 != 0 {
					n.nearestZeroList = append(n.nearestZeroList, middle+1)
				}
				for i := middle; i > 0; i-- {
					n.nearestZeroList = append(n.nearestZeroList, i)
				}
				n.count = 0
			}
			n.nearestZeroList = append(n.nearestZeroList, 0)
		}

		if n.currentHouseNumber != 0 && n.zeroFound {
			n.count++
		}

	}
	for i := 1; i <= n.count; i++ {
		n.nearestZeroList = append(n.nearestZeroList, i)
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
