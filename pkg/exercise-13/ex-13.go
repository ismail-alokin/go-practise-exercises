package ex_13

import (
	"log"
	"os"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_13() {
	question := `
	Create a Go Program that:

	1. Using single function creates a file called info.txt in the current directory. If the file already exists it will truncate it to zero size.

	2. Write the string "The Go gopher is an iconic mascot!" to the file.
	`

	input_output.PrintExerciseTitles(question)

	myfile, err := os.Create("./pkg/exercise-13/info.txt")
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}
	defer myfile.Close()

	myfile.WriteString("The Go gopher is an iconic mascot!")
}
