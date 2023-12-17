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

	numbersMap := make(map[string]int, 10)

	for i := 0; i < 4; i++ {
		scanner.Scan()
		row := scanner.Text()
		a := string(row[0])
		numbersMap[a] = numbersMap[a] + 1
		b := string(row[1])
		numbersMap[b] = numbersMap[b] + 1
		c := string(row[2])
		numbersMap[c] = numbersMap[c] + 1
		d := string(row[3])
		numbersMap[d] = numbersMap[d] + 1
	}

	var points int
	for key, value := range numbersMap {
		if key == "." {
			continue
		}

		if value <= handSize {
			points++
		}
	}

	printInt(points)
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
