package calculator

import (
	Stack "GolangHomeWorkTechnoPark/stack"
	"errors"
	"math"
	"strconv"
	"strings"
)

// Calculator принимает строчку и возвращает результат вычисления этой строчки или ошибку
func Calculator(inputString string) (float64, error) {
	stringToCalc := strings.Split(strings.ReplaceAll(inputString, " ", ""), "")
	stringToCalc = UnionBigNumbers(stringToCalc)
	answer, err := ParseAll(stringToCalc)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

// UnionBigNumbers принимает slice строк, объединяет числа, которые имеют длину больше одного символа,
// в одну строку и возвращает измененный slice строк
func UnionBigNumbers(stringWithoutBigNumbers []string) []string {
	var (
		stringWithBigNumbers []string
		intermediateString   string
	)
	for i := 0; i < len(stringWithoutBigNumbers); i++ {
		_, err := strconv.Atoi(stringWithoutBigNumbers[i])
		if err == nil {
			intermediateString += stringWithoutBigNumbers[i]
			for j := i + 1; j < len(stringWithoutBigNumbers); j++ {
				_, err = strconv.Atoi(stringWithoutBigNumbers[j])
				if err == nil {
					intermediateString += stringWithoutBigNumbers[j]
				} else if stringWithoutBigNumbers[j] == "." {
					intermediateString += stringWithoutBigNumbers[j]
				} else {
					stringWithBigNumbers = append(stringWithBigNumbers, intermediateString)
					i = j - 1
					break
				}
				if j == len(stringWithoutBigNumbers)-1 {
					stringWithoutBigNumbers = append(stringWithBigNumbers, stringWithoutBigNumbers[j])
					stringWithBigNumbers = append(stringWithBigNumbers, intermediateString)
					i = j
				}
			}
		} else {
			stringWithBigNumbers = append(stringWithBigNumbers, stringWithoutBigNumbers[i])
			intermediateString = ""
		}
	}
	if len(intermediateString) == 1 {
		stringWithBigNumbers = append(stringWithBigNumbers, intermediateString)
	}
	return stringWithBigNumbers
}

// ParseAll - основная функция калькулятора, вычисляющая значение переданного среза строчек.
// Возвращает значение выражения или ошибку
func ParseAll(stringToCalc []string) (float64, error) {
	var (
		result float64
	)

	// Цикл разбивает slices на скобки, рекурсивно вызывая ParseAll для каждой из подскобок
	for i := 0; i < len(stringToCalc); i++ {
		if stringToCalc[i] == "(" {
			openIndex := i
			openCount := 1
			closeIndex := i
			for j := i + 1; j < len(stringToCalc); j++ {
				if stringToCalc[j] == "(" {
					openCount++
				}
				if stringToCalc[j] == ")" {
					openCount--
					if openCount == 0 {
						closeIndex = j
						break
					}
				}
			}
			intermediateStringParse, err := ParseAll(stringToCalc[openIndex+1 : closeIndex])
			if err != nil {
				return 0, err
			}
			intermediateString := strconv.FormatFloat(intermediateStringParse, 'E', -1, 64)
			leftPart := stringToCalc[:openIndex]
			leftPart = append(leftPart, intermediateString)
			rightPart := stringToCalc[closeIndex+1:]
			stringToCalc = append(leftPart, rightPart...)
		}
	}

	// Вычисление унарного минуса
	if stringToCalc[0] == "-" {
		stringToCalc[0] = stringToCalc[0] + stringToCalc[1]
		if len(stringToCalc) > 2 {
			stringToCalc = append([]string{stringToCalc[0]}, stringToCalc[2:]...)
		} else {
			stringToCalc = []string{stringToCalc[0]}
		}
	}

	// Вычисления унарного плюса
	if stringToCalc[0] == "+" {
		stringToCalc[0] = stringToCalc[0] + stringToCalc[1]
		if len(stringToCalc) > 2 {
			stringToCalc = append([]string{stringToCalc[0]}, stringToCalc[2:]...)
		} else {
			stringToCalc = []string{stringToCalc[0]}
		}
	}

	// Проверка на ошибки в синтаксисе
	if stringToCalc[0] == "*" || stringToCalc[0] == "/" || len(stringToCalc)%2 == 0 {
		err := errors.New("ошибка синтаксиса")
		return 0, err
	}

	// Цикл проходит по слайсу и вычисляет значения внутри него относительно умножения и деления
	var stack Stack.Stack
	stack.Push(stringToCalc[0])
	for i := 1; i < len(stringToCalc)-1; {
		if stringToCalc[i] == "+" || stringToCalc[i] == "-" {
			if stringToCalc[i+1] == "+" || stringToCalc[i+1] == "-" || stringToCalc[i+1] == "*" || stringToCalc[i+1] == "/" {
				err := errors.New("ошибка синтаксиса")
				return 0, err
			}
			stack.Push(stringToCalc[i])
			stack.Push(stringToCalc[i+1])
		} else {
			pop, _ := stack.Pop()
			firstBinOp, _ := strconv.ParseFloat(pop, 64)
			secondBinOp, _ := strconv.ParseFloat(stringToCalc[i+1], 64)
			if stringToCalc[i] == "*" {
				stack.Push(strconv.FormatFloat(firstBinOp*secondBinOp, 'E', -1, 64))
			} else if stringToCalc[i] == "/" {
				stack.Push(strconv.FormatFloat(firstBinOp/secondBinOp, 'E', -1, 64))
			} else {
				if stringToCalc[i] == "," {
					err := errors.New("ошибка синтаксиса! " +
						"Калькулятор не поддерживает числа с запятой, но поддерживает числа с точкой")
					return 0, err
				}
				err := errors.New("ошибка синтаксиса! Неизвестный символ операции")
				return 0, err
			}
		}
		i += 2
	}
	stringToCalc = stack

	// Цикл, который позволяет пройтись по слайсу и вычислить значения внутри него относительно сложения и вычитания
	for 1 < len(stringToCalc) {
		var intermediateResult float64
		if stringToCalc[1] == "+" {
			firstBinOp, _ := strconv.ParseFloat(stringToCalc[0], 64)
			secondBinOp, _ := strconv.ParseFloat(stringToCalc[2], 64)
			intermediateResult += firstBinOp + secondBinOp
		} else {
			firstBinOp, _ := strconv.ParseFloat(stringToCalc[0], 64)
			secondBinOp, _ := strconv.ParseFloat(stringToCalc[2], 64)
			intermediateResult += firstBinOp - secondBinOp
		}

		rightPart := make([]string, len(stringToCalc)-2)
		copy(rightPart[1:], stringToCalc[3:])
		rightPart[0] = strconv.FormatFloat(intermediateResult, 'E', -1, 64)
		stringToCalc = rightPart
	}
	result, _ = strconv.ParseFloat(stringToCalc[0], 64)
	result = math.Round(result*100000000000000) / 100000000000000
	return result, nil
}
