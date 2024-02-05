package input_output

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PrintExerciseTitles(question string) {
	fmt.Println(question)
	fmt.Print("-------------------\n\n\n")
}

func GetStringArrayOfInputValues(message string, separator string) []string {
	fmt.Print(message)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Println("ERROR:", err)
	}

	inputString := scanner.Text()

	inputArray := strings.Split(inputString, separator)

	filteredInputArray := make([]string, 0, len(inputArray))

	for _, str := range inputArray {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		filteredInputArray = append(filteredInputArray, str)
	}

	return filteredInputArray
}

func GetMultiLineInputStrings(message string) []string {
	fmt.Println(message)

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	return lines
}
