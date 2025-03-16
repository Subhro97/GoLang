// Overflow Conditionals

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func readBalanceFromFile() (float64, error) {
	balance, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Unable to read from file.")
		return 1000, err
	}

	balanceText := string(balance)

	balanceValue, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		fmt.Println("Unable to convert value to float.")

		return 1000, err
	}

	return balanceValue, nil

}

func writeBalanceToFile(balance float64) error {
	balanceText := strconv.FormatFloat(balance, 'f', -1, 64)

	err := os.WriteFile(fileName, []byte(balanceText), 0644) // File Permission Code - Linux Specific

	if err != nil {
		return errors.New("failed to write to file")
	}

	return nil
}

func main() {
	investedAmount, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("Error Occured!")
		fmt.Println("----------")
		fmt.Println("")
		// panic(err)
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do today?")
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

			writeBalanceToFile(investedAmount)

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

			writeBalanceToFile(investedAmount)

			fmt.Printf("Your current Balance: ₹%0.2f\n", investedAmount)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thank You for choosing our bank!")
}
