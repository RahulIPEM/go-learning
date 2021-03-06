package main

import (
	"fmt"
	"os"
)

// Trade struct
type Trade struct {
	Symbol string  // Trade symbol
	Volume int     // Number of trade
	Price  float64 // Price of trade
	Buy    bool    // true if buy trade, false if sell trade
}

// NewTrade will create a new trade will validate the input
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be > 0 (was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return trade, nil
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
	t, err := NewTrade("AMSX", 10, 88.98, true)

	if err != nil {
		fmt.Printf("ERROR: can't create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}
