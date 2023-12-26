package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := makeScanner()

	handSize := readInt(scanner)
	handSize = handSize * 2

	numbers := make([]int, 10)

	for i := 0; i < 4; i++ {
		scanner.Scan()
		row := scanner.Text()

		addValue(row[0], numbers)
		addValue(row[1], numbers)
		addValue(row[2], numbers)
		addValue(row[3], numbers)
	}

	var points int
	for _, value := range numbers {
		if value == 0 {
			continue
		}
		if value <= handSize {
			points++
		}
	}

	printInt(points)
}

func addValue(value uint8, numbers []int) {
	if value == 46 {
		return
	}
	number := int(value - '0')
	numbers[number] += 1
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 20 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printInt(value int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(value))
	writer.Flush()
}
