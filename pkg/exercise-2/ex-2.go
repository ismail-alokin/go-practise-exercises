package ex_2

import (
	"fmt"
	"log"
	"strconv"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_2() {
	question := "Exercise-2: Write a program which takes 2 digits, X,Y as input and generates a 2-dimensional array. \nThe element value in the i-th row and j-th column of the array should be i*j\n"
	message := "Enter comma separated values of X and Y: "
	inputSeperator := ","

	input_output.PrintExerciseTitles(question)
	xy := input_output.GetStringArrayOfInputValues(message, inputSeperator)

	if len(xy) != 2 {
		log.Fatal("ERROR: Exactly 2 values required")
	}

	x, err := strconv.Atoi(xy[0])
	if err != nil {
		log.Fatal("ERROR: X should be an integer")
	}

	y, err := strconv.Atoi(xy[1])
	if err != nil {
		log.Fatal("ERROR: Y should be an integer")
	}

	resultSlice := make([][]int, x)
	for i := range resultSlice {
		resultSlice[i] = make([]int, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			resultSlice[i][j] = i * j
		}
	}

	fmt.Println("The required array: ", resultSlice)
}
