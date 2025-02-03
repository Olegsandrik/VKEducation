package inputOutput

import (
	uniq "GolangHomeWorkTechnoPark/unique"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// Scanner - определяет параметры переданные программе и в соответствии с ними заполняет структуру Options,
// возвращая Options или ошибку
func Scanner() (uniq.Options, error) {
	option := uniq.Options{}
	var err error
	if len(os.Args) < 2 {
		err = errors.New("введите параметры")
		return option, err
	}

	var c = flag.Bool("c", false, "Flag -c")
	var d = flag.Bool("d", false, "Flag -d")
	var u = flag.Bool("u", false, "Flag -u")
	var i = flag.Bool("i", false, "Flag -i")
	var f = flag.Int("f", 0, "Flag -f")
	var s = flag.Int("s", 0, "Flag -s")
	var input = flag.String("input", "", "input file")
	var output = flag.String("output", "", "output file")
	flag.Parse()

	option.C = *c
	option.D = *d
	option.U = *u
	option.I = *i
	option.F = *f
	option.S = *s
	option.Input = *input
	option.Output = *output

	if (option.U && option.D) || (option.C && option.D) || (option.C && option.U) {
		err = errors.New("выберите один из флагов c, d, u")
	}

	if option.Output != "" && option.Input == "" {
		err = errors.New("нельзя передать параметр output без параметра input")
	}

	return option, err
}

// FileWork получает на вход структуру options и на основе ее параметров вызывает функцию Uniq,
// после чего выводит данные либо в консоль, либо в файл, указанный в options
func FileWork(option uniq.Options) {
	inputFile := option.Input
	outputFile := option.Output
	var readFrom io.Reader
	var writeTo io.Writer
	stringsIn := []string{}

	if option.Input != "" { // Либо файлик ввели как параметр
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		readFrom = file
	} else { // Либо данные нужно считать
		readFrom = os.Stdin
	}

	// читаем данные
	scanner := bufio.NewScanner(readFrom)
	for scanner.Scan() {
		stringsIn = append(stringsIn, scanner.Text())
	}

	stringsOut, err2 := uniq.Uniq(stringsIn, option)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	if outputFile != "" { // Либо файлик ввели как параметр
		file2, err3 := os.Create(outputFile)
		if err3 != nil {
			fmt.Println(err3)
			return
		}
		writeTo = file2
		defer file2.Close()
	} else { // Либо выводим в консоль
		writeTo = os.Stdout
	}
	for i := 0; i < len(stringsOut); i++ {
		fmt.Fprintln(writeTo, stringsOut[i])
	}
}
