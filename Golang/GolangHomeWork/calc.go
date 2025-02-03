package main

import (
	myCalc "GolangHomeWorkTechnoPark/calculator"
	myScan "GolangHomeWorkTechnoPark/scan"
	"fmt"
)

func main() {
	scan, err := myScan.Scanner()
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := myCalc.Calculator(scan)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result)
		return
	}
}
