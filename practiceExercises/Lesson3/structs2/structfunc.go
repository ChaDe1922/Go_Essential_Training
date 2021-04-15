//This is a practice in new structs using functions
package main

import (
	"fmt"
	"os"
)

//struct, Trade, represents a trade in stocks
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

//Constructor function
//NewTrade will create a new trade object and will validate the input
//NewTrade receives 4 arguements, the data for the Trade struct
//NewTrade returns a Trade pointer to the created object and an error
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {

	//NewTrade validates by checking each arguement for errors
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >= 0 (was %f)", price)
	}

	//creates a variable, trade, containing a pointer to a trade object (&Trade)
	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil
}

//METHOD, Value, returns the trade value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t, err := NewTrade("MSFT", 10, 99.98, true)

	if err != nil {
		fmt.Printf("error: can't create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}
