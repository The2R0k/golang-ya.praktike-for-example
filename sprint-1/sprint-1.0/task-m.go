package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Ids struct {
	Id    string
	Count int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := strings.Split(scanner.Text(), "")
	idMap := make(map[string]int)

	for i := 0; i < len(inputStr); i++ {
		idMap[inputStr[i]]++
	}

	outputStr := []string{""}
	for i := 0; i < len(inputStr); i++ {
		output := Ids{"", 0}
		for k, v := range idMap {
			if (v > output.Count) || (v == output.Count && k < output.Id) {
				output = Ids{k, v}
			}
		}
		delete(idMap, output.Id)
		for j := 0; j < output.Count; j++ {
			outputStr = append(outputStr, output.Id)
		}
	}
	fmt.Println(strings.Join(outputStr, ""))
}
