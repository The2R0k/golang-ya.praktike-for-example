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
	word1 := strings.Split(scanner.Text(), "")
	scanner.Scan()
	word2 := strings.Split(scanner.Text(), "")

	isAnagramm := "True"

	if len(word1) != len(word2) {
		isAnagramm = "False"
	} else {
		checker := dictionary()
		for i := 0; i < len(word1); i++ {
			checker[word1[i]]++
			checker[word2[i]]--
		}
		for _, value := range checker {
			if value != 1 {
				isAnagramm = "False"
				break
			}
		}
	}

	fmt.Println(isAnagramm)
}

func dictionary() map[string]int {
	checker := map[string]int{
		"a": 1,
		"b": 1,
		"c": 1,
		"d": 1,
		"e": 1,
		"f": 1,
		"g": 1,
		"h": 1,
		"i": 1,
		"j": 1,
		"k": 1,
		"l": 1,
		"m": 1,
		"n": 1,
		"o": 1,
		"p": 1,
		"q": 1,
		"r": 1,
		"s": 1,
		"t": 1,
		"u": 1,
		"v": 1,
		"w": 1,
		"x": 1,
		"y": 1,
		"z": 1,
		"а": 1,
		"б": 1,
		"в": 1,
		"г": 1,
		"д": 1,
		"е": 1,
		"ё": 1,
		"ж": 1,
		"з": 1,
		"и": 1,
		"й": 1,
		"к": 1,
		"л": 1,
		"м": 1,
		"н": 1,
		"о": 1,
		"п": 1,
		"р": 1,
		"с": 1,
		"т": 1,
		"у": 1,
		"ф": 1,
		"х": 1,
		"ц": 1,
		"ч": 1,
		"ш": 1,
		"щ": 1,
		"ъ": 1,
		"ы": 1,
		"ь": 1,
		"э": 1,
		"ю": 1,
		"я": 1,
	}
	return checker
}
