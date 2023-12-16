package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	daysLen := readInt(scanner)
	if daysLen == 1 {
		fmt.Println(1)
		return
	}
	temperatures := readArray(scanner)

	weatherRandomness := getWeatherRandomness(temperatures, daysLen)

	fmt.Println(weatherRandomness)
}

func getWeatherRandomness(temperatures []int, daysLen int) int {
	var randomnessIndex int

	for index, temperature := range temperatures {
		var afterTemp, beforeTemp int
		afterIndex := index + 1
		beforeIndex := index - 1

		if index == 0 {
			afterTemp = temperatures[afterIndex]
			if afterTemp < temperature {
				randomnessIndex++
			}
			continue
		}

		if index == daysLen-1 {
			beforeTemp = temperatures[beforeIndex]
			if beforeTemp < temperature {
				randomnessIndex++
			}
			continue
		}

		beforeTemp = temperatures[beforeIndex]
		afterTemp = temperatures[afterIndex]
		if beforeTemp < temperature && afterTemp < temperature {
			randomnessIndex++
		}
	}

	return randomnessIndex
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

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}
