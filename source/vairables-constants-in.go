package main

import "fmt"

func main() {
	var someVariableNameHere int = 1
	fmt.Println(someVariableNameHere)

	var (
		anotherVariableNameHere = 4
		yetAnotherVariableHere  = 34.45
	)
	fmt.Println(anotherVariableNameHere)
	fmt.Println(yetAnotherVariableHere)

	var username string = "your-username-here"
	fmt.Println(username)

	const pi float64 = 3.14159
	fmt.Println(pi)

	const ageLimit int = 18
	fmt.Println(ageLimit)

	const addressTag string = "@"
	fmt.Println(addressTag + username)

	const hashTag string = "#"
	fmt.Println(hashTag + "Golang")

}
