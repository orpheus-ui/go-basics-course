package main

import "fmt"

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate (default is 2.5): ")

	earningBeforeTax, earningAfterTax, earningRatio := profitCalculator(revenue, expenses, taxRate)

	formattedRatio := fmt.Sprintf(`Ratio: %.2f`, earningRatio)

	fmt.Print("Earning Before Tax: ", earningBeforeTax, " | ", "Profit: ", earningAfterTax, " | ", formattedRatio)
}

func getUserInput(title string) float64 {
	var userInput float64
	fmt.Print(title)
	fmt.Scan(&userInput)
	return userInput
}

func profitCalculator(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - taxRate/100)
	er := ebt / eat
	return ebt, eat, er
}
