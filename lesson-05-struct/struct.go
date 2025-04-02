package main

import (
	"fmt"

	"example.com/struct/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthday (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Print(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "123456")

	admin.OutputUserDetails()

	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
