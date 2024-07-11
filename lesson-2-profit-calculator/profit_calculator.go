package main

import (
	"fmt"
	"os"
)

// Goal
// 1) Validate User Input
// 2) Store Calculated results into file

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate (default is 2.5): ")

	earningBeforeTax, earningAfterTax, earningRatio := profitCalculator(revenue, expenses, taxRate)

	writeToFile(earningBeforeTax, "ebt.txt")
	writeToFile(earningAfterTax, "profit.txt")
	writeToFile(earningRatio, "ratio.txt")

	formattedRatio := fmt.Sprintf(`Ratio: %.2f`, earningRatio)

	fmt.Print("Earning Before Tax: ", earningBeforeTax, " | ", "Profit: ", earningAfterTax, " | ", formattedRatio)
	fmt.Println("New values added to respective files.")
}

func getUserInput(title string) float64 {
	var userInput float64
	for {
		fmt.Print(title)
		fmt.Scan(&userInput)
		if userInput <= 0 {
			fmt.Println("Error: Value should be above 0!")
			continue
		} else {
			break
		}
	}

	return userInput
}

func profitCalculator(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - taxRate/100)
	er := ebt / eat
	return ebt, eat, er
}
func writeToFile(i float64, str string) {
	value := fmt.Sprint(i)
	os.WriteFile(str, []byte(value), 0644)
}
