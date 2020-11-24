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
	scanner.Scan()

	// "0"
	replacer := strings.NewReplacer(" 0", "")
	if strings.HasPrefix(scanner.Text(), "0") {
		fmt.Println(replacer.Replace(strings.TrimPrefix(scanner.Text(), "0")))
	} else {
		fmt.Println(replacer.Replace(scanner.Text()))
	}
}
