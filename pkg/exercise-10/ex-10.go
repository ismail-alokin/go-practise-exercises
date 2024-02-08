package ex_10

import (
	"fmt"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_10() {
	question := "There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.\n"

	input_output.PrintExerciseTitles(question)

	mySlice := []float64{1.2, 5.6}

	mySlice = append(mySlice, 6)

	a := float64(10)
	mySlice[0] = a

	mySlice = append(mySlice, 10.10)

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
	fmt.Println("Errors in the program have been fixed and is runnning properly")

}

// mySlice := []float64{1.2, 5.6}

// mySlice = append(mySlice, 6)

// a := float64(10)
// mySlice[0] = a

// mySlice[3] = 10.10

// mySlice = append(mySlice, a)

// fmt.Println(mySlice)

// }
