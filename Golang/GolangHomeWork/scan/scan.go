package scan

import (
	"errors"
	"os"
)

// Scanner - функция сканирующая аргументы из входной строки. Ничего не принимает
// и возвращает входную строку или ошибку
func Scanner() (string, error) {
	var inputString string
	if len(os.Args) < 2 {
		err := errors.New("введите числовое выражение")
		return "", err
	}
	inputString = os.Args[1]
	stringToCalc := inputString
	return stringToCalc, nil
}
