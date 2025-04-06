package main

import "fmt"

func main() {
	var age int = 32

	agePointer := &age

	fmt.Println("Age:", *agePointer) // de-references the value from the address

	adultYears(agePointer)

	fmt.Println(age)
}

func adultYears(age *int) {
	*age = *age - 18 // Reassigning the value in memory
}
