package ex_1

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

const C, H int = 50, 30

func Ex_1() {
	question := "Exercise-1: Write a program that calculates and prints the value according to the given formula:\nQ = Square root of [(2 * C * D)/H],\nwhere C = 50, H = 30\n"
	message := "Enter comma separated values of D: "
	inputSeperator := ","

	input_output.PrintExerciseTitles(question)
	valuesOfD := input_output.GetStringArrayOfInputValues(message, inputSeperator)

	var valuesOfQ []string

	C := 50
	H := 30

	for _, D := range valuesOfD {
		D = strings.TrimSpace(D)
		if D == "" {
			continue
		}
		D, err := strconv.Atoi(D)
		if err != nil {
			log.Fatal("ERROR: All values should be integers")
		}

		product := (2 * C * D) / H
		Q := math.Sqrt(float64(product))

		valuesOfQ = append(valuesOfQ, strconv.FormatFloat(Q, 'f', 0, 64))
	}

	answer := strings.Join(valuesOfQ, ",")
	fmt.Println("The required values of Q are: ", answer)
}
