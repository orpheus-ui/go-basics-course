package main

import (
	"errors"
	"fmt"
	"os"
)

// Goal
// 1) Validate User Input
// 2) Store Calculated results into file

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate (default is 2.5): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningBeforeTax, earningAfterTax, earningRatio := profitCalculator(revenue, expenses, taxRate)

	formattedRatio := fmt.Sprintf(`Ratio: %.2f`, earningRatio)

	fmt.Print("Earning Before Tax: ", earningBeforeTax, " | ", "Profit: ", earningAfterTax, " | ", formattedRatio)

	writeToFile(earningBeforeTax, earningAfterTax, earningRatio)

}

func getUserInput(title string) (float64, error) {
	var userInput float64
	fmt.Print(title)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("value most be positive number")
	}
	return userInput, nil
}

// Tutor wants to teach "errors" package but I found below function much simpler and works better in this case
// since it loops and ask the value until it's right

// func getUserInputWithError(title string) float64 {
// 	var userInput float64
// 	for {
// 		fmt.Print(title)
// 		fmt.Scan(&userInput)
// 		if userInput <= 0 {
// 			println("Error: Value should be positive!")
// 			continue
// 		} else {
// 			break
// 		}
// 	}
// 	return userInput
// }

func profitCalculator(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := ebt * (1 - taxRate/100)
	er := ebt / eat
	return ebt, eat, er
}

func writeToFile(ebt, profit, ratio float64) {
	data := fmt.Sprintf("Earning before Task: %v\nProfit: %v\nRatio: %v", ebt, profit, ratio)
	os.WriteFile("data.txt", []byte(data), 0644)
}
