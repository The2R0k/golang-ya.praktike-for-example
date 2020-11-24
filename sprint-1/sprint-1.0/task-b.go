package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ids struct {
	Id    string
	Count int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	id := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	numbK, _ := strconv.Atoi(scanner.Text())
	idMap := make(map[string]int)

	for i := 0; i < len(id); i++ {
		idMap[id[i]]++
	}

	for i := 0; i < numbK; i++ {
		output := Ids{"", 0}
		for k, v := range idMap {
			if (v > output.Count) || (v == output.Count && k < output.Id) {
				output = Ids{k, v}
			}
		}
		delete(idMap, output.Id)
		fmt.Print(output.Id, " ")
	}
}
