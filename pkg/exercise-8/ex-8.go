package ex_8

import (
	"fmt"
	"unicode"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_8() {
	question := "Write a program that accepts a sentence and calculate the number of letters and digits.\n"
	message := "Input the string"

	input_output.PrintExerciseTitles(question)
	inputStr := input_output.GetAString(message)

	lettersCount := 0
	digitsCount := 0

	for _, ltr := range inputStr {
		if unicode.IsLetter(ltr) {
			lettersCount++
			continue
		}

		if unicode.IsDigit(ltr) {
			digitsCount++
			continue
		}
	}

	fmt.Printf("LETTERS: %d \nDIGITS: %d\n", lettersCount, digitsCount)

}
