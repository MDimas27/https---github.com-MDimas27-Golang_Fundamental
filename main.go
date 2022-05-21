package main

import (
	"errors"
	"fmt"
)

func main() {

	//result, err := calculate(10, 2, "+")
	//result, err := calculate(10, 2, "-")
	//result, err := calculate(10, 2, "*")
	//result, err := calculate(10, 2, "/")
	//result, err := calculate(10, 2, "=")

	result, err := calculate(10, 2, "m")
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(result)


}

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("Operasi Tidak dikenali")
	}

	return result, errorResult
}
