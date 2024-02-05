package ex_3

import (
	"fmt"
	"slices"
	"strings"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_3() {
	question := "Write a program that accepts a comma separated sequence of words as input and prints the words in a comma-separated sequence after sorting them alphabetically.\n"
	message := "Enter comma separated strings: "
	inputSeperator := ","

	input_output.PrintExerciseTitles(question)
	splitStrings := input_output.GetStringArrayOfInputValues(message, inputSeperator)

	slices.Sort(splitStrings)
	sortedString := strings.Join(splitStrings, ",")

	fmt.Println("Sorted string: ", sortedString)
}
