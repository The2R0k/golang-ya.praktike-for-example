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
	str := strings.Split(strings.ToLower(scanner.Text()), "")
	i := 0
	j := len(str) - 1
	checker := dictionary()
	isPolyndrom := true
	for i < j && isPolyndrom {
		if checker[str[i]] {
			if checker[str[j]] {
				if str[i] == str[j] {
					i++
					j--
				} else {
					isPolyndrom = false
				}
			} else {
				j--
			}
		} else {
			i++
		}
	}
	if isPolyndrom {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}

func dictionary() map[string]bool {
	checker := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"e": true,
		"f": true,
		"g": true,
		"h": true,
		"i": true,
		"j": true,
		"k": true,
		"l": true,
		"m": true,
		"n": true,
		"o": true,
		"p": true,
		"q": true,
		"r": true,
		"s": true,
		"t": true,
		"u": true,
		"v": true,
		"w": true,
		"x": true,
		"y": true,
		"z": true,
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
		"0": true,
	}
	return checker
}
