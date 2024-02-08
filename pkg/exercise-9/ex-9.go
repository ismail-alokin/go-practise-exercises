package ex_9

import (
	"fmt"
	"strings"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_9() {
	question := `A website requires the users to input a username and password to register. Write a program to check the validity of password input by users. The following are the criteria for checking the password:
	1. At least 1 letter between [a-z]
	2. At least 1 number between [0-9]
	3. At least 1 letter between [A-Z]
	4. At least 1 character from [$#@]
	5. Minimum length of transaction password: 6
	6. Maximum length of transaction password: 12
	`
	message := "Enter passwords to verify: (separated by comma) "
	inputSeperator := ","

	input_output.PrintExerciseTitles(question)
	passwords := input_output.GetStringArrayOfInputValues(message, inputSeperator)

	validPasswords := make([]string, 0)
	for _, password := range passwords {
		if CheckValidPassword(password) {
			validPasswords = append(validPasswords, password)
		}
	}
	if len(validPasswords) == 0 {
		fmt.Println("No valid passwords provided")
	} else {
		result := strings.Join(validPasswords, inputSeperator)
		fmt.Println("Valid Passwords:", result)
	}

}
