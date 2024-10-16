package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	//seed our random number generator
	rand.Seed(time.Now().UnixNano())
	//print welocme message
	color.Yellow("The sleeping Barber Problem")
	color.Yellow("---------------------------")
	//create channels if needed
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)
	//create the barbershop
	shop := barbershop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientChan:      clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")
	//add barbers
	shop.addBarber("Frank")
	shop.addBarber("Tom")
	shop.addBarber("Dan")
	shop.addBarber("Fred")
	shop.addBarber("Don")
	shop.addBarber("Fred3")
	shop.addBarber("Don4")
	shop.addBarber("Fred5")
	shop.addBarber("Don6")

	//start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closShopForDay()
		closed <- true

	}()
	//add clients
	i := 1
	go func() {
		for {
			//get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barber is closed
	<-closed
}
