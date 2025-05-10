package main

import (
	"fmt"
	"time"
)

type Result struct {
	prices []Price
}

type Price struct {
	ProcessNumber int64
	Price         int64
}

func main() {
	// Create a struct to hold the results
	var getPrices = &Result{
		prices: []Price{},
	}

	// Create a timeout channel to handle the timeout after 3 seconds
	timeout := time.After(3 * time.Second)

	// Start the long process in a goroutine
	go longProcess1(getPrices)
	go longProcess2(getPrices)

	select {
	case <-timeout:
		fmt.Println("Timeout occurred with prices: ", getPrices.prices)
	}

	// Wait for the results or timeout
	maxPrice := getMaxPrice(getPrices)
	fmt.Println(fmt.Sprintf("Process %d is winner with prince %d", maxPrice.ProcessNumber, maxPrice.Price))

	time.Sleep(time.Second * 10)
}

func getMaxPrice(results *Result) *Price {
	maxPrice := &Price{
		ProcessNumber: 0,
		Price:         0,
	}
	for _, p := range results.prices {
		if p.Price > maxPrice.Price {
			maxPrice = &p
		}
	}
	return maxPrice
}

func loggingProcess(processNumber int64, res int64) {
	fmt.Println(fmt.Sprintf(`Get process %d result %d`, processNumber, res))
}

func longProcess1(getPrices *Result) {
	fmt.Println("Processing data 1....")

	time.Sleep(2 * time.Second)
	getPrices.prices = append(getPrices.prices, Price{
		ProcessNumber: 1,
		Price:         9000,
	})
	go loggingProcess(1, 9000)
}

func longProcess2(getPrices *Result) {
	fmt.Println("Processing data 2....")

	time.Sleep(2 * time.Second)
	getPrices.prices = append(getPrices.prices, Price{
		ProcessNumber: 2,
		Price:         9000,
	})
	go loggingProcess(2, 9000)
}
