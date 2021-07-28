package main

import "fmt"

func main() {
	// example for int data type
	var basicSalary int = 10000
	fmt.Println(basicSalary)

	const ageLimit = 18
	fmt.Println(ageLimit)

	variableSalary :=10000
	fmt.Println(variableSalary)

	// example for float data type
	const pi float64 = 3.14159
	fmt.Println(pi)

	const piAgain = 3.14159
	fmt.Println(piAgain)

	increaseSalaryNextYearBy := 3.14159
	fmt.Println(increaseSalaryNextYearBy)

	// example for string data type
	var username string = "your-username-here"
	fmt.Println(username)

	const addressTag = "@"
	fmt.Println(addressTag + username)

	hashTag := "#"
	fmt.Println(hashTag + "Golang")

	// short way to declare and store any data type
	performanceBonus, increaseSalaryNextYearBy, username := 2000, 0.6, "my-user-name"
	fmt.Println(performanceBonus)
	fmt.Println(increaseSalaryNextYearBy)
	fmt.Println(username)

	const piCorrection, ageLimitRevised, addressTagForDeletedAccount = 3.1416, 16, "."
	fmt.Println(piCorrection)
	fmt.Println(ageLimitRevised)
	fmt.Println(addressTagForDeletedAccount+username)

}
