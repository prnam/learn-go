package main

import "fmt"

func main()  {

	// for loop example
	fmt.Println("Iterate using for loop")
	for  item := 1; item <= 10 ; item++ {
		fmt.Println(item)
	}

	// achieving while loop in Go
	fmt.Println("Achieving while loop in Go")
	count := 1
	for count <= 10{
		fmt.Println(count)
		count++
	}
}
