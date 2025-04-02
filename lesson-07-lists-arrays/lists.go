package main

import "fmt"

func main() {

	var productNames [4]string = [4]string{"Book"}

	prices := [4]float64{10.99, 9.99, 45.99, 20}

	fmt.Println(prices)

	productNames[3] = "Carpet"

	fmt.Println(productNames)
	fmt.Println(prices[2], prices[0])

	// From index:1 to index:3 (index:to index (this index number exclude from list)
	featuredPrices := prices[1:3] // [9.99, 45.99]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3] // [10.99 9.99 45.99]
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:]            //[9.99 45.99 20] | Here we're starting to ommit left value (capacity 3)
	featuredPrices[0] = 199.99             //Change index0 Value
	higlightedPrices := featuredPrices[:1] //[199.99]
	fmt.Println(higlightedPrices)
	fmt.Println(prices)

	fmt.Println(len(higlightedPrices), cap(higlightedPrices)) // Length & Capacity
	/*
		The capacity only counts towards the end of the original array,
		but omits any elements that might have been filtered out before.
	*/

	higlightedPrices = higlightedPrices[:3]
	fmt.Println(higlightedPrices)
	fmt.Println(len(higlightedPrices), cap(higlightedPrices))

}
