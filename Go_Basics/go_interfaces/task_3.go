package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Vehicle interfeysi
type Vehicle interface {
	GetId() int
	Move() string
	GetPrice() float64
}

// ID uchun global hisoblagich
var idCounter int32
var mu sync.Mutex

func generateID() int {
	return int(atomic.AddInt32(&idCounter, 1))
}

// Car struct
type Car struct {
	model string
	id    int
	price float64
}

func (c Car) Move() string {
	return "The car " + c.model + " is moving."
}

func (c Car) GetId() int {
	return c.id
}

func (c Car) GetPrice() float64 {
	return c.price
}

// Bike struct
type Bike struct {
	brand string
	id    int
	price float64
}

func (b Bike) Move() string {
	return "The bike " + b.brand + " is moving."
}

func (b Bike) GetId() int {
	return b.id
}

func (b Bike) GetPrice() float64 {
	return b.price
}

func main() {
	var vehicles []Vehicle

	// Avtomobil va velosipedni mos ravishda avtomatik id va narx bilan to'ldirish
	vehicles = append(vehicles, Car{model: "Tesla", id: generateID(), price: 60000})
	vehicles = append(vehicles, Bike{brand: "Yamaha", id: generateID(), price: 15000})

	for _, vehicle := range vehicles {
		fmt.Println(vehicle.Move())
		fmt.Printf("ID: %d\n", vehicle.GetId())
		fmt.Printf("Price: $%.2f\n", vehicle.GetPrice())
	}
}
