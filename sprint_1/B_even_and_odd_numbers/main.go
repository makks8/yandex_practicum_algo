package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	playerWon  = "WIN"
	playerLose = "FAIL"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var a, b, c int
	scanner.Scan()
	line := scanner.Text()
	values := strings.Split(line, " ")
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var mainParityIsEven bool
	var numberIsEven bool
	a, _ = strconv.Atoi(values[0])
	mainParityIsEven = isEven(a)

	b, _ = strconv.Atoi(values[1])
	numberIsEven = isEven(b)
	if mainParityIsEven && !numberIsEven || !mainParityIsEven && numberIsEven {
		writer.WriteString(playerLose)
		return
	}

	c, _ = strconv.Atoi(values[2])
	numberIsEven = isEven(c)

	if mainParityIsEven && !numberIsEven || !mainParityIsEven && numberIsEven {
		writer.WriteString(playerLose)
		return
	}

	writer.WriteString(playerWon)
	writer.WriteString("\n")
}

func isEven(num int) bool {
	return num%2 == 0
}
