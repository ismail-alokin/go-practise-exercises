package ex_7

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_7() {
	question := "Write a program, which will find all such numbers between A and B (both included) such that each digit of the number is an even number. The numbers obtained should be printed in a comma-separated sequence on a single line.\n"
	message := "Enter A and B (separated by comma): "
	inputSeperator := ","

	input_output.PrintExerciseTitles(question)
	limits := input_output.GetStringArrayOfInputValues(message, inputSeperator)

	if len(limits) != 2 {
		log.Fatal("Enter lower and upper limits only")
	}

	lowerLimit, err := strconv.Atoi(limits[0])
	if err != nil {
		log.Fatal("Enter only integers")
	}

	upperLimit, err := strconv.Atoi(limits[1])
	if err != nil {
		log.Fatal("Enter only integers")
	}

	if lowerLimit >= upperLimit {
		log.Fatal("Enter lower and upper limit correctly")
	}

	numbersWithAllDigitsEven := make([]string, 0)
	for num := lowerLimit; num <= upperLimit; num++ {
		digits := ParseDigitsOfInteger(num)

		allDigitsEven := true
		for _, digit := range digits {
			if digit%2 != 0 {
				allDigitsEven = false
				break
			}
		}

		if allDigitsEven {
			numbersWithAllDigitsEven = append(numbersWithAllDigitsEven, strconv.Itoa(num))
		}
	}
	result := strings.Join(numbersWithAllDigitsEven, inputSeperator)
	fmt.Println("These are the numbers with all even digits:", result)

}
