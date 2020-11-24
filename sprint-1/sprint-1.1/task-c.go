package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := strings.Split(scanner.Text(), "")

	dictionary := make(map[string]int)

	length := len(str)
	maxLength := 0
	curLength := 0
	for i := 0; i < length; i++ {
		if dictionary[str[i]] == 0 {
			curLength++
		} else {
			if maxLength < curLength {
				maxLength = curLength
			}
			curLength = i - (dictionary[str[i]] - 1)
		}
		dictionary[str[i]] = i + 1
		fmt.Println(i, maxLength, curLength)
	}
	if maxLength < curLength {
		maxLength = curLength
	}
	fmt.Println(maxLength)
}
