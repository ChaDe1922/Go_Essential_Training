//This is an example of maps. A map is a data structure where keys point to values
//In go, the keys must be of the same type and the values must be of the same type

package main

import (
	"fmt"
)

func main() {

	//we define a map from [strings] to floats
	//set the initial values inside curly braces (remember the trailing comma)
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	//Number of Items in a map
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MSFT"])

	// Get zero value if not found
	fmt.Println(stocks["TSLA"]) //0

	// Use double value contexts to see if it is found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}
	//set
	stocks["TSLA"] = 322.12

	//Delete from a map
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	//Single value "for". This will bring us only the keys
	for key := range stocks {
		fmt.Println(key)
	}

	//Double value for. this will bring us the key and the value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}

}
