package lists

import "fmt"

func main() {
	prices := []float64{10.99}
	fmt.Println(prices[0])

	updatedPrices := append(prices, 9.99) // append actually returns a new slice
	fmt.Println(updatedPrices, prices)

	prices = append(prices, 8.99) // We Should reassign the variable to the first slice returned by append
	fmt.Println(prices)

	prices = prices[1:] // Remove the first element from the slice
	fmt.Println(prices)

	discountedPrices := []float64{7.99, 6.99, 5.99}
	prices = append(prices, discountedPrices...) // append the discounted prices to the slice
	fmt.Println(prices)
}
