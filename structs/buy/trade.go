package main

import (
	"fmt"
)

// Trade struct
type Trade struct {
	Symbol string  // Trade symbol
	Volume int     // Number of trade
	Price  float64 // Price of trade
	Buy    bool    // true if buy trade, false if sell trade
}

// Value returns the trade value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t := Trade{
		Symbol: "AMAZ",
		Volume: 20,
		Price:  99.75,
		Buy:    true,
	}

	fmt.Println(t.Value())
}
