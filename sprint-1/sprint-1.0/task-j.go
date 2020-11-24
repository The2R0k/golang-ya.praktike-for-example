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
		data[numbs[i]] = !data[numbs[i]]
	}
	for key, value := range data {
		if value {
			fmt.Println(key)
			break
		}
	}
}
