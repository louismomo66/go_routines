package main

import (
	"time"

	"github.com/fatih/color"
)

type barbershop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *barbershop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)
		for {
			if len(shop.ClientChan) == 0 {
				color.Yellow("There is nothing to do so %s takes a nap.", barber)
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				shop.cutHair(barber, client)
			} else {
				//shop is closed, so send the barber home and close
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *barbershop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (shop *barbershop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *barbershop) closShopForDay() {
	color.Cyan("Closing shop fpr the day.")
	close(shop.ClientChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)
	color.Green("------------------------------")
	color.Green("The shop is closed for the day")
}

func (shop *barbershop) addClient(client string) {
	color.Green("*** %s arrival!", client)
	if shop.Open {
		select {
		case shop.ClientChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves!", client)
	}
}
