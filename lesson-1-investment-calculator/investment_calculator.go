package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var clientName string
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Your Name: ")
	fmt.Scan(&clientName)

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedName := fmt.Sprintf("Hey %v\n", clientName)
	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Real Profit (Inflation): %.2f", futureRealValue)

	fmt.Println("-----------")
	fmt.Println(formattedName, formattedFV, formattedRFV)
	fmt.Println("-----------")

	// fmt.Println("--------------")
	// fmt.Println("Hey ", clientName)
	// fmt.Println("Your Future Value: ", futureValue)
	// fmt.Println("Real Value (w/ Tax): ", futureRealValue)
	// fmt.Printf("------------\nHey %v\nFuture Value: %.2f\nReal Profit (Inflation): %.2f\n------------", clientName, futureValue, futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
