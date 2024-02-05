package main

import (
	"fmt"
	"log"

	ex_1 "github.com/ismail-alokin/go-practise-exercises/pkg/exercise-1"
	ex_2 "github.com/ismail-alokin/go-practise-exercises/pkg/exercise-2"
	ex_3 "github.com/ismail-alokin/go-practise-exercises/pkg/exercise-3"
	ex_4 "github.com/ismail-alokin/go-practise-exercises/pkg/exercise-4"
	ex_5 "github.com/ismail-alokin/go-practise-exercises/pkg/exercise-5"
	ex_6 "github.com/ismail-alokin/go-practise-exercises/pkg/exercise-6"
	ex_7 "github.com/ismail-alokin/go-practise-exercises/pkg/exercise-7"
)

var exNum int

func main() {
	log.SetFlags(log.Lshortfile)

	fmt.Println(`Enter the exercise number you want to run: (To run exercise-1, type 1 and press ENTER)`)
	_, err := fmt.Scan(&exNum)

	if err != nil {
		log.Println("ERROR:", err)
	}

	fmt.Printf("You chose Exercise-%d\n\n", exNum)

	switch exNum {

	case 1:
		ex_1.Ex_1()

	case 2:
		ex_2.Ex_2()

	case 3:
		ex_3.Ex_3()

	case 4:
		ex_4.Ex_4()

	case 5:
		ex_5.Ex_5()

	case 6:
		ex_6.Ex_6()

	case 7:
		ex_7.Ex_7()

	default:
		fmt.Println("-------- Exercise not found --------")

	}
}
