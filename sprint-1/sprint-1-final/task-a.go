package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doMathOp(stack []string, operation string) []string {
	n := len(stack) - 1
	b, _ := strconv.Atoi(stack[n])
	a, _ := strconv.Atoi(stack[n-1])
	stack = stack[:n-1]
	switch operation {
	case "+":
		stack = append(stack, strconv.Itoa(a+b))
	case "-":
		stack = append(stack, strconv.Itoa(a-b))
	case "*":
		stack = append(stack, strconv.Itoa(a*b))
	case "/":
		stack = append(stack, strconv.Itoa(a/b))
	}
	return stack
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")

	var stack []string

	for i := range input {
		_, err := strconv.Atoi(input[i])
		if err != nil {
			stack = doMathOp(stack, input[i])
		} else {
			stack = append(stack, input[i])
		}
	}
	fmt.Println(stack[0])

}
