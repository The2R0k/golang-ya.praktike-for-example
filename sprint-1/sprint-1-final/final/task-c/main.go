/*\
 |	Для данной задачи было важным реализовать Очередь через Стэк.
 |  Суть в том, что стэк опирается на идиолигию FILO, в то время как
 |	очередь - FIFO. Однако, в условии задачи не сказано что необходимо
 | 	использовать только один стэк. Второй стэк нужен для "Разворота"
 |	данных в стэке, за счет чего и достигается нужная очередность эллементов
 |	Алгоритм работает за O(N), и использует O(1) дополнительной памяти
 |	(т.к. совокупность размеров двух стэков не превышает значение N)
 |  Номер удачной послыки: 38627406 (с правильной реализацией стэка)
\*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func doGet(stack1 []string, stack2 []string) ([]string, []string) {
	lastStack2 := len(stack2) - 1
	if lastStack2 >= 0 {
		fmt.Println(stack2[lastStack2])
		stack2 = stack2[:lastStack2]
	} else {
		if len(stack1) > 0 {
			stack1, stack2 = swopStacks(stack1)
			stack1, stack2 = doGet(stack1, stack2)
		} else {
			fmt.Println("error")
		}
	}
	return stack1, stack2
}

func swopStacks(stack []string) (stack1 []string, stack2 []string) {
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	stack2 = stack
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandsNumb, _ := strconv.Atoi(scanner.Text())

	var stack1, stack2 []string

	for i := 0; i < commandsNumb; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		switch input[0] {
		case "put":
			stack1 = append(stack1, input[1])
		case "get":
			stack1, stack2 = doGet(stack1, stack2)
		case "get_size":
			fmt.Println(len(stack1) + len(stack2))
		}
	}
}
