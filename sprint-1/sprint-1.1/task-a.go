package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	visitedOptionalCourses := make(map[string]bool)
	for i := 0; i < count; i++ {
		scanner.Scan()
		scanner.Text()
		if visitedOptionalCourses[scanner.Text()] == false {
			visitedOptionalCourses[scanner.Text()] = true
			fmt.Println(scanner.Text())
		}
	}
}
