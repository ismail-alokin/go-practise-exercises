package ex_6

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_6() {
	question := "Write a program which accepts a sequence of comma separated 4 digit binary numbers as its input and then check whether they are divisible by 5 or not. The numbers that are divisible by 5 are to be printed in a comma separated sequence.\n"
	message := "Enter comma separated binary numbers: "
	inputSeperator := ","

	input_output.PrintExerciseTitles(question)
	binaryNumbers := input_output.GetStringArrayOfInputValues(message, inputSeperator)

	numbersDivisibleByFive := make([]string, 0, len(binaryNumbers))

	for _, num := range binaryNumbers {
		num, err := strconv.ParseInt(num, 2, 64)
		if err != nil {
			log.Fatal("ERROR: All numbers should be binary")
		}

		if num%5 == 0 {
			numbersDivisibleByFive = append(numbersDivisibleByFive, strconv.FormatInt(num, 2))
		}
	}

	result := strings.Join(numbersDivisibleByFive, inputSeperator)
	fmt.Println("Numbers divisible by 5 are: ", result)
}
