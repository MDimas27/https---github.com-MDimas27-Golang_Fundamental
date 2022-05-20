package main

import "fmt"

func main() {

	// hitung rata2
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var total int
	for _, score := range(scores) {
		total = total + score
	}

	length := len(scores)
	average := total / length

	fmt.Println(average)


}
