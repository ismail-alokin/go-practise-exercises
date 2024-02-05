package ex_5

import (
	"fmt"
	"slices"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_5() {
	question := "Write a program that accepts a sequence of whitespace separated words as input and prints the words after removing all duplicate words and sorting them alphanumerically\n"
	message := "Enter space separated strings: "
	inputSeperator := " "

	input_output.PrintExerciseTitles(question)
	splitStrings := input_output.GetStringArrayOfInputValues(message, inputSeperator)

	distinctElementsArray := make([]string, 0, len(splitStrings))

	for _, str := range splitStrings {
		if slices.Contains(distinctElementsArray, str) {
			continue
		}
		distinctElementsArray = append(distinctElementsArray, str)
	}

	slices.Sort(distinctElementsArray)
	fmt.Println("Sorted distinct array: ", distinctElementsArray)
}
