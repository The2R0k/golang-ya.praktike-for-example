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
	x, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	fmt.Println(strconv.FormatInt(x, 2))
}
