package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const hunger = 3

var philosophers = []string{"Plato", "Socrates", "Aristole", "pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var orderOfFinish []string
var orderMutex sync.Mutex

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()
	fmt.Println(philosopher, "is seated.")
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "is hungry.")
		time.Sleep(sleepTime)

		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left.\n", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right.\n", philosopher)

		fmt.Println(philosopher, "has both forks,and is eating.")
		time.Sleep(eatTime)

		fmt.Println(philosopher, "is thinking.")
		time.Sleep(thinkTime)

		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right.\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left.\n", philosopher)
		time.Sleep(sleepTime)
	}
	fmt.Println(philosopher, "is satisfied")

	time.Sleep(sleepTime)
	fmt.Println(philosopher, "hasleft the table.")
	orderMutex.Lock()
	orderOfFinish = append(orderOfFinish, philosopher)
	orderMutex.Unlock()
}
func main() {
	fmt.Println("The dining philosophers problem")
	fmt.Println("-------------------------------")
	wg.Add(len(philosophers))

	forkLeft := &sync.Mutex{}
	for i := 0; i < len(philosophers); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(philosophers[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	wg.Wait()
	fmt.Println("The table is empty and the philosophers finished in the order: ")
	fmt.Println("-------------------------------")
	fmt.Printf("Oderfinished: %s\n", strings.Join(orderOfFinish, ","))

}
