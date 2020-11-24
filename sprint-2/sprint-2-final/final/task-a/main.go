/*\
 | 	Изначально, я реализовал алгоритм при котором были выбраны самый большой,
 | 	и самый малый датацентры, а в них помещалось по одной фотографии. Неудачно.
 | 	Остовалось очень много позиций. Тогда я попробовал алгоритм где выбирались
 | 	самые большие датацентры. Тоже неочень. НО! После этого я решил объеженить
 | 	две стратегии, а именно: выбирать большие датаценетры, и складывать туда
 | 	по 1-ой фотографии. Таким образом я смогу вместить наибольшее количество фото.
 | 	Работает примерно так : [1 2 3 4 5] > [1 2 3 (4) (5)] > [1 2 3 (3) (4)] > sort >
 |  > [1 2 3 3 4] > [1 2 3 (3) (4)]  > [1 2 3 (2) (3)] > sort  > [1 2 2 3 3] > ... >
 |	> [0 0 1 1 1] > [1 1 1] > ...
 | 	Для наиболее эффективной работы была написана сортировка при которой перемещаются
 |  только два последних эллемента. Тип сортировки - пузырьком (самое смешное, но она
 |  здесь и самая быстрая, и самая простая в написании).
 |	Алгоритм работае за O(N) и использует O(1) дополнительной памяти
 |  Номер удачной посылки: 42301831
\*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func removeFilledDC(dataCapacity *[]int) {
	for len(*dataCapacity) > 0 && (*dataCapacity)[0] == 0 {
		*dataCapacity = (*dataCapacity)[1:]
	}
}

func fastSort(dc *[]int, elem int) {
	length := len(*dc)
	if length > 1 {
		for i := 1; i < elem; i++ {
			if (*dc)[elem-i] < (*dc)[elem-i-1] {
				(*dc)[elem-i-1], (*dc)[elem-i] = (*dc)[elem-i], (*dc)[elem-i-1]
			} else {
				break
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dataCenterCount, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	dataCapacityStrings := strings.Split(scanner.Text(), " ")
	dataCapacity := make([]int, dataCenterCount)
	for ; dataCenterCount > 0; dataCenterCount-- {
		dataCapacity[dataCenterCount-1], _ = strconv.Atoi(dataCapacityStrings[dataCenterCount-1])
	}

	sort.Ints(dataCapacity)
	removeFilledDC(&dataCapacity)

	i := 0
	for len(dataCapacity) > 1 {
		dataCapacity[len(dataCapacity)-1]--
		dataCapacity[len(dataCapacity)-2]--
		fastSort(&dataCapacity, len(dataCapacity)-1)
		fastSort(&dataCapacity, len(dataCapacity))
		removeFilledDC(&dataCapacity)
		i++
	}
	fmt.Println(i)
}
