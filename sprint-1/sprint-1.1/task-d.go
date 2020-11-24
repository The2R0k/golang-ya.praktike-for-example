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
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	collumn, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	row, _ := strconv.Atoi(scanner.Text())

	array := make([][]string, collumn)

	for i := 0; i < collumn; i++ {
		scanner.Scan()
		array[i] = strings.Split(scanner.Text(), " ")
	}

	scanner.Scan()
	posX, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	posY, _ := strconv.Atoi(scanner.Text())

	var output []string

	switch posX {
	case 0:
		if collumn != 1 {
			output = append(output, array[posX+1][posY])
		}
	case collumn - 1:
		output = append(output, array[posX-1][posY])
	default:
		output = append(output, array[posX-1][posY], array[posX+1][posY])
	}

	switch posY {
	case 0:
		if row != 1 {
			output = append(output, array[posX][posY+1])
		}
	case row - 1:
		output = append(output, array[posX][posY-1])
	default:
		output = append(output, array[posX][posY+1], array[posX][posY-1])
	}

	outputInt := make([]int, len(output))

	for i := range output {
		outputInt[i], _ = strconv.Atoi(output[i])
	}
	sort.Ints(outputInt)
	for i := range output {
		output[i] = strconv.Itoa(outputInt[i])
	}

	fmt.Println(strings.Join(output, " "))
}
