/*\
 |	В этой задаче я использовал алгоритм поиска циклов "Floyd's Tortoise and Hare".
 |	Суть алгоритма сводится к тому, что у нас есть 2 указателя на список, которые
 |	двигаются с разной скорость (i=1 для черепахи - Tortoise и i=2 для зайца - Hare).
 |  В какой то момент времени (из-за разницы скоростей) оба указателя встречаются
 |	на одной позиции, либо быстрый указатель дохоит до конца списка.
 |	Я не нашел сложность этого алгоритма, однако, я бы оценил её как O(NLogN) т.к.
 |	на больших длинах списка мы можем блуждать по циклу довольно долго.
 |	Простанственная сложность алгоритма - O(1)
 |  Номер удачной послыки: 41788068
\*/

package main

// HasCycle is used to found cycle in TListItem struct
func HasCycle(head *TListItem) string {
	answer := doHasCycle(head, head.nextElement)
	return answer
}

func doHasCycle(tortoise *TListItem, hare *TListItem) string {
	if hare == nil || hare.nextElement == nil {
		return "False"
	}
	if tortoise == hare {
		return "True"
	}
	return doHasCycle(tortoise.nextElement, hare.nextElement.nextElement)
}
