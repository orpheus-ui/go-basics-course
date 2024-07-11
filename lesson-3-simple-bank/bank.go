package main

import (
	"fmt"

	fo "example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fo.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic("Can't continue, Sorry!")
	}
	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()
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
			fo.WriteFloatToFile(accountBalance, accountBalanceFile)

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
			fo.WriteFloatToFile(accountBalance, accountBalanceFile)
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
