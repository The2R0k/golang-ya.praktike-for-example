/*\
 |	Для данной задачи было важным реализовать тип работы с данными "Стэк".
 |	Начитавшись умных форумов, я узнал что правильная реализация стэка на Go
 |	через массив, с постоянным переопределением онного.
 |	Алгоритм работает за O(N), и использует O(1) дополнительной памяти
 |  Номер удачной послыки: 34217009
\*/

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

/* P.S. В Тесте 23 ошибка связаная с описанием задачи
 * "Операция деления целочисленная. То есть, например, 12 5 / будет 2."
 * | Входной файл
 * | 2 5 - 4 /
 * | Вывод программы
 * | 0
 * | Правильный ответ
 * |-1
 *	i=0 :  2-5 -> -3
 *	i=1 : -3/4 -> -0.75 (0 при целочисленном делении)
 */
