package main

import "fmt"

func main() {

	title := "Golang the best language"
	fmt.Println("============= MENAMPILKAN ELEMEN DENGAN INDEX GENAP =============")
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index : ", index, "letter : ", string(letter))
		}
	}

	fmt.Println("============= MENAMPILKAN ELEMEN DENGAN HURUF VOCAL SAJA =============")
	for index, letter := range title {
		letterString := string(letter)
		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
			fmt.Println("index : ", index, "letter : ", string(letter))
		}

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index : ", index, "letter : ", string(letter))
	}
	}


}
