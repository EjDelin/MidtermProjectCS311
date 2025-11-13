package Debugging

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Amount float64
}

func processOrder(o Order, wg *sync.WaitGroup, total *float64) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Processing order %d\n", o.ID)

	*total += o.Amount

	if o.Amount < 0 {
		fmt.Printf("Warning: Negative amount for order %d\n", o.ID)
	}
}

func main() {
	orders := []Order{
		{1, 100},
		{2, -50}, // wrong order
		{3, 200},
	}

	var total float64
	var wg sync.WaitGroup

	for _, order := range orders {
		wg.Add(1)
		go processOrder(order, &wg, &total)
	}

	wg.Wait()
	fmt.Printf("\nExpected total: 250 | Got: %.2f\n", total)
}
