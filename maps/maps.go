package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMAZ": 102.6,
		"GOOG": 2009.32,
		"WIPR": 332.43,
	}

	// number of items
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["GOOG"])

	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("Not found")
	} else {
		fmt.Println(value)
	}

	stocks["TSLA"] = 322.12
	fmt.Println(stocks["TSLA"])

	delete(stocks, "TSLA")

	for key := range stocks {
		fmt.Println(key)
	}

	for key, val := range stocks {
		fmt.Printf("%s -> %.2f", key, val)
	}
}
