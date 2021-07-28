package main

import "fmt"

func main() {
	whichNumber := 20

	fmt.Println(factorial(whichNumber))
}

func factorial(whichNumber int) int {

	if whichNumber == 0 {
		return 1
	}

	return whichNumber * factorial(whichNumber-1)

}
