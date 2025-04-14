package main

import (
	"fmt"

	"github.com/structs/user"
)

type str string

func (s str) log() {
	fmt.Println(s)
}

func main() {
	userFirstName := getUserData("Enter your First Name: ")
	userLastName := getUserData("Enter your Last Name: ")
	userBirthDate := getUserData("Enter your BirthDate: ")

	userData, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}
	userData.OutputUserDetails()
	userData.ClearUserDetails()
	userData.OutputUserDetails()

	const name str = "Subhro"

	name.log()
}

func getUserData(outputStr string) string {
	fmt.Print(outputStr)
	var value string
	fmt.Scanln(&value)
	return value
}
