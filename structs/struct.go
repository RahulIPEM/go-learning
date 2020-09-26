package main

import (
	"fmt"
)

// Trade struct
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of Stocks
	Price  float64 // Price of stocks
	Buy    bool    // true if buy trade, false if sell trade
}

func main() {
	t1 := Trade{"MSFT", 10, 999.34, true}

	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "AMAZ",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}

	fmt.Println(t3)
}
