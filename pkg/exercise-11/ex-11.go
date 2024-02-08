package ex_11

import (
	"fmt"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_11() {
	question := `Consider the following array declaration: nums := [...]int{30, -1, -6, 90, -6}

	Write a Go program that counts the number of positive even numbers in the array.
	`
	input_output.PrintExerciseTitles(question)

	nums := [...]int{30, -1, -6, 90, -6}

	resultCount := 0
	for _, num := range nums {
		if num < 0 {
			continue
		}

		if num%2 == 0 {
			resultCount++
		}
	}

	fmt.Printf("There are %d positive even numbers\n", resultCount)
}
