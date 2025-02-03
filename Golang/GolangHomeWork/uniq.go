package main

import (
	inOut "GolangHomeWorkTechnoPark/inputOutput"
	"fmt"
)

func main() {
	// Запуск:
	// go run uniq.go -c -i -f=0 -s=0 -input=input.txt -output=output.txt
	// cat input.txt | go run uniq.go -c -i -f=1 -s=1 -output=output.txt
	option, err := inOut.Scanner()
	if err != nil {
		fmt.Println(err)
		return
	}
	inOut.FileWork(option)

}
