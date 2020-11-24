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
	length, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	numbs := strings.Split(scanner.Text(), " ")
	data := make(map[string]bool)
	for i := 0; i < length; i++ {
		if !data[numbs[i]] {
			data[numbs[i]] = true
		} else {
			fmt.Println(numbs[i])
			break
		}
	}
}
