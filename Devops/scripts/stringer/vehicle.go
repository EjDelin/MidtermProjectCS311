package main

import "fmt"

//go:generate stringer -type=Vehicle
type Vehicle int

const (
	Car Vehicle = iota
	Bike
	Truck
	Bus
	Airplane
	Tuktuk
	Train
)

func main() {
	//vehicleNames := []string{"Car", "Bike", "Truck", "Bus", "Airplane", "Tuktuk", "Train"}
	vehicles := []Vehicle{Car, Bike, Truck, Bus, Airplane, Tuktuk, Train}

	for _, v := range vehicles {
		fmt.Println(v)
	}
}
