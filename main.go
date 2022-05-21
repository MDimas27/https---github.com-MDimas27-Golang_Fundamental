package main

import "fmt"

func main() {
	// scores := []int {10, 5, 8, 9, 7}
	// total := sum(scores)
	scores := []int {10, 5, 8, 9, 7}
	
	total := sum(scores)
	fmt.Println(total)

}

func sum(numbers []int) int{

	var total int
	for _, number := range numbers {
		total = total + number
	}

	return total


}
