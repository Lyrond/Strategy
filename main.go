package main

import "fmt"

var (
	ticketPrice = 69.69
	strategies  = []Strategy{
		FirstTimePerform{},
		CalendarPrice{},
		NewAlbumPrice{},
	}
)

type Calculator struct {
	Strategy
}

func (calc *Calculator) SetStrategy(strategy Strategy) {
	calc.Strategy = strategy
}

type CalendarPrice struct {
}

func (c CalendarPrice) PriceCalculation(ticketPrice float64) {
	discount := 10
	totalPrice := ticketPrice - (ticketPrice*float64(discount))/100
	fmt.Printf("Common price: [%f], but for these dates discount is [%d]. So the total price for ticket is [%f]", ticketPrice, discount, totalPrice)
	fmt.Println("")
}

type FirstTimePerform struct {
}

func (c FirstTimePerform) PriceCalculation(ticketPrice float64) {
	discount := 15
	totalPrice := ticketPrice - (ticketPrice*float64(discount))/100
	fmt.Printf("Common price: [%f], but for the first band's perform the discount is [%d], the total price for ticket is [%f]", ticketPrice, discount, totalPrice)
	fmt.Println("")
}

type NewAlbumPrice struct {
}

func (c NewAlbumPrice) PriceCalculation(ticketPrice float64) {
	discount := 20
	totalPrice := ticketPrice - (ticketPrice*float64(discount))/100
	fmt.Printf("Common price: [%f], as the band will perform their new album, your dicount is [%d], the total price for ticket is [%f]", ticketPrice, discount, totalPrice)
	fmt.Println("")
}

type Strategy interface {
	PriceCalculation(ticketPrice float64)
}

func main() {
	calculator := Calculator{}

	var event string
	fmt.Scanln(&event)
	switch event {
	case "Album":
		calculator.SetStrategy(NewAlbumPrice{})
		calculator.PriceCalculation(ticketPrice)

	case "fti":
		calculator.SetStrategy(FirstTimePerform{})
		calculator.PriceCalculation(ticketPrice)
	case "calendar":
		calculator.SetStrategy(CalendarPrice{})
		calculator.PriceCalculation(ticketPrice)
	}
}
