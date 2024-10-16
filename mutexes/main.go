package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	//vsriable for bank balance
	var bankBalance int
	var balance sync.Mutex
	//print out the starting value
	fmt.Printf("inittial account balance: %d.00", bankBalance)
	fmt.Println()

	//define weekly revenue
	incomes := []Income{
		{Source: "main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time", Amount: 50},
		{Source: "Investments", Amount: 100},
	}
	wg.Add(len(incomes))
	//loop trough
	for _, income := range incomes {

		go func(income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earned %d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(income)
	}
	wg.Wait()

	fmt.Printf("Final bank balance: %d.00", bankBalance)
}
