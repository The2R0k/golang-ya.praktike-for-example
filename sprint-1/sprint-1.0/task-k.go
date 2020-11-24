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
	x, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	str := strings.Split(strconv.FormatInt(x, 2), "")
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == "1" {
			count++
		}
	}
	fmt.Println(count)
}
