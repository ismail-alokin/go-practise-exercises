package ex_12

import (
	"fmt"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_12() {
	question := `Consider the following map declaration: m := map[int]bool{"1": true, 2: false, 3: false}
	1. The above map declaration has an error. Solve the error!
	
	2. Delete a key:value pair from the map.
	
	3. Iterate over the map and print out each key and value.
	`

	input_output.PrintExerciseTitles(question)

	m := map[int]bool{1: true, 2: false, 3: false}
	fmt.Println("Corrected the map declaration. Current value of m: ", m)

	fmt.Println("Deleting a key value pair: (2: false)")
	delete(m, 2)
	fmt.Println("Deleted key value pair - 2: false. Current value of m: ", m)

	fmt.Println("Iterating over map and printing out existing keys and values")
	for key, value := range m {
		fmt.Println("\t", key, value)
	}

}
