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

	scanner.Scan()
	collumn, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	row, _ := strconv.Atoi(scanner.Text())

	array := make([][]string, row)
	for i := 0; i < 0; i++ {
		array[i] = make([]string, collumn)
	}

	for i := 0; i < collumn; i++ {
		scanner.Scan()
		str := strings.Split(scanner.Text(), " ")
		for j := 0; j < row; j++ {
			array[j][i] = str[j]
		}
	}
	fmt.Println(array)
}
