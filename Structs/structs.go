package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Enter your First Name: ")
	userLastName := getUserData("Enter your Last Name: ")
	userBirthDate := getUserData("Enter your BirthDate: ")

	var userData user

	userData = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	outputUserDetails(userData)
}

func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(outputStr string) string {
	fmt.Print(outputStr)
	var value string
	fmt.Scan(&value)
	return value
}
