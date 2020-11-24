package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//	string 1: length
	scanner.Scan()
	length, _ := strconv.Atoi(scanner.Text())

	// string 2: number X
	scanner.Scan()
	numbXString := strings.Split(scanner.Text(), " ")

	numbX := 0
	for i := 0; i < length; i++ {
		numbX = numbX * 10
		x, _ := strconv.Atoi(numbXString[i])
		numbX = numbX + x
	}

	// string 3: number K
	scanner.Scan()
	numbK, _ := strconv.Atoi(scanner.Text())

	// X + K and output
	output := ""
	outputNumb := numbX + numbK
	for outputNumb > 0 {
		x := outputNumb % 10
		output = strings.Join([]string{strconv.Itoa(x), output}, " ")
		outputNumb = (outputNumb - x) / 10
	}

	fmt.Println(output)
}
