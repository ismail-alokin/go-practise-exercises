package ex_15

import (
	"fmt"
	"sync"

	input_output "github.com/ismail-alokin/go-practise-exercises/pkg/utils"
)

func Ex_15() {
	question := `Protect the account's balance from simultaneously writing using a mutex.`
	input_output.PrintExerciseTitles(question)

	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}

func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b += n
	m.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b -= n
	m.Unlock()
	wg.Done()
}
