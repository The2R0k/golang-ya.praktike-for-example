/*\
 | 	При разработке программы я опирался на этот алгоритм. https://neerc.ifmo.ru/wiki/index.php?title=%D0%A6%D0%B5%D0%BB%D0%BE%D1%87%D0%B8%D1%81%D0%BB%D0%B5%D0%BD%D0%BD%D1%8B%D0%B9_%D0%B4%D0%B2%D0%BE%D0%B8%D1%87%D0%BD%D1%8B%D0%B9_%D0%BF%D0%BE%D0%B8%D1%81%D0%BA
 | 	Если вкратце - мы делим массив на 3 части (|/|\|/| или |\|/|\|) и ищем элемент в каждой из 3-ех частей.
 | 	А когда я подумал, то понял что не нужно искать в каждой из трех частей,
 | 	а массив и вовсе можно поделить на 2 части. Отсортированная, и не отстортированная!
 | 	Зная какой элемент мы ищем, мы можем сказать 100% может ли он быть в упорядоченной
 | 	части массива или нет. На этом и базируется логика алгоритма.
 | 	Т.е. Берем последовательность [5 6 7 8 9 0 1 2 3 4],
 |  делаем 3 замера того, возрастает он или нет [(5 6) 7 8 (9 0) 1 2 (3 4)]
 |  на этих замерах строим догадку, растет или уменьшаяется массив (направление сортировки неизвестноф)
 |  Берем узловой (центральный) элемент [5 6 7 8 (9) 0 1 2 3 4]
 |	Этот элемент является левой границей для правого массива, и правой для левого.
 | 	Смотрим кто из них больше, а кто меньше (Можно предполагать, что в массиве только уникальные элементы)
 |  Если "направление" сортировки половины совпадает с глобальным, мы считаем что эта последовательность отсортирована
 |  Если не совпадает, то она "хаотична".
 |  Сравниваем, входит ли элемент в границы отсортированной последовательности. Если да, он может быть только там.
 |  Если нет, значит он в "хаотичном" отрезке. Повторяем пока не найдем элемент.
 |	Алгоритм работае за O(logN) и использует O(1) дополнительной памяти (коппирование массива же не в счет? да?)
 |  Номер удачной посылки: 42460983
\*/

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
	length, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))
	scanner.Scan()
	find, _ := strconv.Atoi(strings.Trim(scanner.Text(), " "))

	scanner.Scan()
	arrayStr := strings.Split(scanner.Text(), " ")
	array := make([]int, length)
	length--
	for ; length >= 0; length-- {
		array[length], _ = strconv.Atoi(arrayStr[length])
	}
	if len(array) == 1 {
		if find == array[0] {
			fmt.Println(0)
		} else {
			fmt.Println(-1)
		}

	} else {
		binarySearch(array, find)
	}
}

func binarySearch(array []int, find int) {
	// Инициализация указателей
	lPtr := 0
	rPtr := len(array) - 1
	curPos := rPtr / 2
	result := 0

	// Определение направления сортировки по 3-ем замерам
	// 0-1 - убывающая 		(>)
	// 2-3 - возрастающая 	(<)

	upper := 0
	if array[lPtr] < array[lPtr+1] {
		upper++
	}
	if array[curPos] < array[curPos+1] {
		upper++
	}
	if array[rPtr-1] < array[rPtr] {
		upper++
	}

	for {
		if len(array) == 0 {
			result = -1
			break
		}

		curPos = len(array) / 2
		rPtr = len(array) - 1
		lPtr = 0

		if find == array[curPos] {
			result += curPos
			break
		}

		// Определение направления движения
		if upper > 1 {
			// Здесь код для возрастающих последовательностей (<)
			if array[curPos] < find {
				if find <= array[rPtr] || array[lPtr] <= array[curPos] {
					// Наше число в отсортированном / неотсортированном масиве справа
					result += curPos + 1
					array = array[curPos+1:]
				} else {
					// Нашего числа нет справа, а массив слева не отсортирован
					array = array[:curPos]
				}
			} else {
				if array[lPtr] <= find || array[curPos] <= array[rPtr] {
					// Наше число в отсортированном / неотсортированном масиве слева
					array = array[:curPos]
				} else {
					// Нашего числа нет слева, а массив справа не отсортирован
					result += curPos + 1
					array = array[curPos+1:]
				}
			}
		} else {
			// Здесь код для убывающих последовательностей (>)
			if find > array[curPos] {
				if array[lPtr] >= find || array[curPos] >= array[lPtr] {
					// Наше число в отсортированном / неотсортированном масиве слева
					array = array[:curPos]
				} else {
					// Нашего числа нет слева, а массив справа не отсортирован
					result += curPos + 1
					array = array[curPos+1:]
				}
			} else {
				if find >= array[rPtr] || array[rPtr] >= array[curPos] {
					// Наше число в отсортированном / неотсортированном масиве справа
					result += curPos + 1
					array = array[curPos+1:]
				} else {
					// Нашего числа нет справа, а массив слева не отсортирован
					array = array[:curPos]
				}
			}
		}
	}
	fmt.Println(result)
}
