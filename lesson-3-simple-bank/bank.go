package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// Function that helps us READ a file using "os" Package
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

// Function that helps us WRITE a file using "os" Package
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic("Can't continue, Sorry!")
	}
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your Balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("How much to Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("------------")
				fmt.Println("Invalid Amount. Should be greater than 0")
				fmt.Println("------------")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("------------")
			fmt.Println("New Amount: ", accountBalance)
			fmt.Println("------------")
			writeBalanceToFile(accountBalance)

		} else if choice == 3 {
			fmt.Print("Withdrawal Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("------------")
				fmt.Println("Invalid Amount. Should be greater than 0")
				fmt.Println("------------")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("------------")
				fmt.Println("Invalid Amount. More Amount of your values")
				fmt.Println("------------")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("------------")
			fmt.Println("New Amount:", accountBalance)
			fmt.Println("------------")
			writeBalanceToFile(accountBalance)
		} else if choice == 4 {
			fmt.Print("Thank you!")
			break
		} else {
			fmt.Println("------------")
			fmt.Println("Choose ONLY between 1-4")
			fmt.Println("------------")
		}
	}
}
