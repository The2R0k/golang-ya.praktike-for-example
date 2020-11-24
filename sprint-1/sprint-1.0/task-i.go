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
	numb1 := strings.Split(scanner.Text(), "")
	scanner.Scan()
	numb2 := strings.Split(scanner.Text(), "")

	lengthM := 0
	length1 := len(numb1)
	length2 := len(numb2)
	if length1 > length2 {
		lengthM = len(numb1) + 1
	} else {
		lengthM = len(numb2) + 1
	}
	numb := make([]string, lengthM)

	c := "0"
	for i := lengthM - 1; i >= 0; i-- {
		a := "0"
		if length1 > 0 {
			length1--
			a = numb1[length1]
		}
		b := "0"
		if length2 > 0 {
			length2--
			b = numb2[length2]
		}
		count := 0
		array := [3]string{a, b, c}
		for char := range array {
			if array[char] == "1" {
				count++
			}
		}
		switch count {
		case 0: // 0 0 0
			numb[i] = "0"
		case 1: // 0 0 1 / 0 1 0 / 1 0 0
			numb[i] = "1"
			c = "0"
		case 3: // 1 1 1
			numb[i] = "1"
			c = "1"
		default: // 0 1 1 / 1 0 1 / 1 1 0
			numb[i] = "0"
			c = "1"
		}

	}
	i := 0
	if numb[0] == "0" {
		i++
	}
	for i < lengthM {
		fmt.Print(numb[i])
		i++
	}
}
