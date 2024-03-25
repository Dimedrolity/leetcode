// Package task_tinkoff_sre
package task_tinkoff_sre

import (
	"strconv"
	"strings"
)

/*
ЗАДАЧА 1
Дана строка из латинских заглавных букв. Необходимо заменить все повторы одинаковых подряд идущих букв на букву + цифру. Одиночные буквы заменять не надо.

Примеры

encode("AAAABBBC") => "A4B3C"
encode("AAAABBBCAAA") => "A4B3CA3"
*/

func encode(s string) string {
	bytes := []byte(s)

	b := strings.Builder{}
	count := 1 // текущее кол-во повторяющихся символов
	for i := 1; i < len(bytes); i++ {
		if bytes[i] != bytes[i-1] {
			b.WriteByte(bytes[i-1])

			if count > 1 {
				b.WriteString(strconv.Itoa(count))
			}
			count = 1
		} else {
			count++
		}
	}

	b.WriteByte(bytes[len(bytes)-1])
	if count > 1 {
		b.WriteString(strconv.Itoa(count))
	}

	return b.String()
}
