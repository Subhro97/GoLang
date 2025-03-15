// Overflow Conditionals

package main

import "fmt"

func main() {
	var investedAmount float64 = 1000

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

			fmt.Printf("Your current Balance: ₹%0.2f\n", investedAmount)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thank You for choosing our bank!")
}
