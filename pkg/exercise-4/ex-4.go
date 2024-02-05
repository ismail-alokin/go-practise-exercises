package ex_4

import (
	"fmt"
	"strings"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_4() {
	question := "Write a program that accepts a sequence of lines as input and prints the lines after making all characters in the sentence capitalized.\n"
	message := "Enter lines of string. Press enter to input next line. Press enter to complete input. "

	input_output.PrintExerciseTitles(question)
	lines := input_output.GetMultiLineInputStrings(message)

	var resultSlice = make([]string, len(lines))

	for _, line := range lines {
		upperCaseSlice := strings.ToUpper(line)
		resultSlice = append(resultSlice, upperCaseSlice)
	}

	resultString := strings.Join(resultSlice, "\n")
	fmt.Println("Strings converted to upper case:\n", resultString)

}
