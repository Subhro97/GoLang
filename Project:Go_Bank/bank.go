// Overflow Conditionals

package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/go-bank/fileops"
)

const fileName = "balance.txt"

func main() {
	investedAmount, err := fileops.ReadFloatFromFile(fileName)

	if err != nil {
		fmt.Println("Error Occured!")
		fmt.Println("----------")
		fmt.Println("")
		// panic(err)
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do today?")
		fmt.Printf("Reach us 24x7 at %v\n", randomdata.PhoneNumber())
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your Balance: ₹%0.2f\n", investedAmount)
		} else if choice == 2 {
			var depositAmt float64
			fmt.Print("Your Deposit: ₹")
			fmt.Scan(&depositAmt)

			if depositAmt <= 0 {
				fmt.Println("Please enter an Amount gretaer than 0.0")
				continue
			}

			investedAmount += depositAmt

			fileops.WriteFloatToFile(fileName, investedAmount)

			fmt.Printf("Your current Balance: ₹%0.2f\n", investedAmount)
		} else if choice == 3 {
			var withdrawAmt float64
			fmt.Print("Your Withdraw: ₹")
			fmt.Scan(&withdrawAmt)

			if withdrawAmt <= 0 {
				fmt.Println("Please enter an Amount gretaer than 0.0")
				continue
			}

			if withdrawAmt > investedAmount {
				fmt.Println("Cannot withdraw Amount gretaer than Balance.")
				continue
			}

			investedAmount -= withdrawAmt

			fileops.WriteFloatToFile(fileName, investedAmount)

			fmt.Printf("Your current Balance: ₹%0.2f\n", investedAmount)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thank You for choosing our bank!")
}
