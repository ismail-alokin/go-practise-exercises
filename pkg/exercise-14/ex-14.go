package ex_14

import (
	"fmt"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_14() {
	question := `
	Consider the following variable declaration:x, y := 5.5, 8.8
	
	Write a function that swaps the values of x and y. After calling the function x will be 8.8 and y will 5.5	
	`
	input_output.PrintExerciseTitles(question)

	x, y := 5.5, 8.8

	fmt.Println("Before swap: (x, y) => ", x, y)
	swap(&x, &y)
	fmt.Println("After swap: (x, y) => ", x, y)
}

func swap(x *float64, y *float64) {
	*x, *y = *y, *x
}
