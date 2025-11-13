package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Amount float64
}

func processOrder(o Order, wg *sync.WaitGroup, total *float64 /*mu *sync.Mutex*/) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Processing order %d\n", o.ID)

	// BUG 1: concurrent write to shared variable (race condition)
	*total += o.Amount

	// BUG 2: logic bug, forgot to check for negative orders
	if o.Amount < 0 {
		fmt.Printf("Warning: Negative amount detected for order %d\n", o.ID)
	}

}

func main() {
	orders := []Order{
		{ID: 1, Amount: 100},
		{ID: 2, Amount: -50}, // invalid order
		{ID: 3, Amount: 200},
	}

	var total float64
	var wg sync.WaitGroup
	//var mu sync.Mutex

	for _, order := range orders {
		wg.Add(1)
		go processOrder(order, &wg, &total /*&mu*/)
	}

	wg.Wait()
	fmt.Printf("\nExpected total: 250 | Calculated total: %.2f\n", total)
}
