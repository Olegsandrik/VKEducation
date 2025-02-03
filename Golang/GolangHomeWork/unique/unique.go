package unique

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Options - Структура флагов
type Options struct {
	C      bool // false
	D      bool
	U      bool
	I      bool
	F      int // 0
	S      int
	Input  string // ""
	Output string
}

// Uniq - Функция уникализации, принимает на вход срез строк и параметрами, возвращает уникальный срез исходя из параметров или ошибку
func Uniq(stringsIn []string, option Options) ([]string, error) {
	var (
		stringsOut        []string
		stringsOutWithC   []string // Соблюдаем CDU или его отсутствие и все остальные флаги
		stringsOutWithD   []string
		stringsOutWithU   []string
		stringsOutWithout []string
	)

	numChars := option.S
	numFields := option.F

	counterUniq := 0
	currentString := stringsIn[0]
	counterUniq++
	stringsOutWithout = append(stringsOut, currentString)
	stringsOutWithC = append(stringsOut, currentString)
	for i := 1; i < len(stringsIn); i++ {
		currentPointer := 0
		stringPointer := 0
		// Выполним условие флага f
		if numFields > 0 {
			countFieldCurrent := 0
			for j := 0; j < len(currentString)-1; j++ {
				if currentString[j] != ' ' && currentString[j+1] == ' ' {
					countFieldCurrent++
					if countFieldCurrent == numFields {
						currentPointer = j + 1 // указывает на байт!
						break
					}
				}
			}

			countFieldString := 0
			for k := 0; k < len(stringsIn[i])-1; k++ {
				if stringsIn[i][k] != ' ' && stringsIn[i][k+1] == ' ' {
					countFieldString++
					if countFieldString == numFields {
						stringPointer = k + 1 // указывает на байт!
						break
					}
				}
			}
		}

		// корректировка, если -f и -s одновременно
		if numChars != 0 && numFields != 0 {
			// Не учитывая пробел-разделитель после последнего поля. - из условия задачки
			// Я так понял, что нужно выкинуть лишь 1 пробел, а остальные значимыми считаются
			currentPointer++
			stringPointer++
			if utf8.RuneCountInString(currentString) < 2 || utf8.RuneCountInString(stringsIn[i]) < 2 {
				err := errors.New("ввели значения флагов -s и -f, превышающие длину введенных строк")
				return []string{}, err
			}
		}
		if utf8.RuneCountInString(currentString) < numChars || utf8.RuneCountInString(stringsIn[i]) < numChars {
			err := errors.New("ввели значения флага -s, превышающие длину введенных строк")
			return []string{}, err
		}

		if option.I {
			// Не учитывать регистр
			// сначала обрезаем field, а потом обрезаем char
			if strings.ToLower(string(([]rune(currentString[currentPointer:]))[numChars:])) == strings.ToLower(string(([]rune(stringsIn[i][stringPointer:]))[numChars:])) {
				counterUniq++
				continue
			} else {
				currentString = stringsIn[i]
				stringsOutWithout = append(stringsOutWithout, currentString)
				stringsOutWithC = append(stringsOutWithC, strconv.Itoa(counterUniq))
				stringsOutWithC = append(stringsOutWithC, currentString)
				counterUniq = 1
			}
		} else {
			// Учитывать регистр
			if string(([]rune(currentString[currentPointer:]))[numChars:]) == (string(([]rune(stringsIn[i][stringPointer:]))[numChars:])) {
				counterUniq++
				continue
			} else {
				currentString = stringsIn[i]
				stringsOutWithout = append(stringsOutWithout, currentString)
				stringsOutWithC = append(stringsOutWithC, strconv.Itoa(counterUniq))
				stringsOutWithC = append(stringsOutWithC, currentString)
				counterUniq = 1
			}
		}

	}

	// На основе массива для флага С вычисляем массивы для D и U
	stringsOutWithC = append(stringsOutWithC, strconv.Itoa(counterUniq))
	for i := 1; i < len(stringsOutWithC); {
		currentCount, _ := strconv.Atoi(stringsOutWithC[i])
		if currentCount > 1 {
			stringsOutWithD = append(stringsOutWithD, stringsOutWithC[i-1])
		} else {
			stringsOutWithU = append(stringsOutWithU, stringsOutWithC[i-1])
		}
		i += 2
	}

	// теперь у нас есть 4 массива для каждого из вариантов CDU или его отсутствие и все остальные флаги

	if option.C {
		// Подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.
		result := []string{}
		for i := 1; i < len(stringsOutWithC); {
			result = append(result, stringsOutWithC[i]+" "+stringsOutWithC[i-1])
			i += 2
		}
		stringsOut = result
	} else if option.D {
		// Вывести только те строки, которые повторились во входных данных.
		stringsOut = stringsOutWithD
	} else if option.U {
		// Вывести только те строки, которые не повторились во входных данных.
		stringsOut = stringsOutWithU
	} else {
		stringsOut = stringsOutWithout
	}

	return stringsOut, nil
}
