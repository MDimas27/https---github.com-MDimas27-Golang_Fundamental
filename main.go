package main

import "fmt"

func main() {

	// Basic Function
	// 1. Parameter/Input
	// 2. Proses
	// 3. Output
	sentence := printMyResult("saya sedang")
	fmt.Println(sentence)

}

func printMyResult(sentence string) string {
	newSentence := sentence + " belajar golang"
	return newSentence
}