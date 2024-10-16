package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	PizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)
		rmd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rmd < 5 {
			pizzaFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rmd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rmd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			PizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}

	return &PizzaOrder{
		PizzaNumber: pizzaNumber,
	}
}
func pizzeria(pizzamaker *Producer) {
	//keep track of which pizza we are making
	var i = 0
	//run forever or until we receive a quit notification
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.PizzaNumber
			select {
			case pizzamaker.data <- *currentPizza:

			case quitChan := <-pizzamaker.quit:
				close(pizzamaker.data)
				close(quitChan)
				return
			}
		}
	}

}
func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//printout the message
	color.Cyan("The Pizzeriia is open for business!")
	color.Cyan("-----------------------------------")
	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	//run the producer in the background
	go pizzeria(pizzaJob)

	//create and run the consumer
	for i := range pizzaJob.data {
		if i.PizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.PizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Dane making pizzas... ")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Failed to close")
			}
		}
	}
	//print out the ending message
	color.Cyan("-------------------")
	color.Cyan("Done for the day.")
	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzaFailed, total)
	switch {
	case pizzaFailed > 9:
		color.Red("It was an awful day...")
	case pizzaFailed >= 6:
		color.Red("It was not a very good day..")
	case pizzaFailed >= 4:
		color.Yellow("It was an okay day...")
	case pizzaFailed >= 2:
		color.Yellow("It was a pretty good day!")
	default:
		color.Green("It was a good day...")
	}

}
