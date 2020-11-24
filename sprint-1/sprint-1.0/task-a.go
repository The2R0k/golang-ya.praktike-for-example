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
	str := strings.Split(scanner.Text(), " ")
	array := make([]int, 4)
	array[0], _ = strconv.Atoi(str[0])
	array[1], _ = strconv.Atoi(str[1])
	array[2], _ = strconv.Atoi(str[2])
	array[3], _ = strconv.Atoi(str[3])

	fmt.Println(array[0]*array[1]*array[1] + array[0]*array[2] + array[3])
}
