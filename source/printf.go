package main

import "fmt"

func main()  {
	// Printf allows to change it to the format specified by you
	// example to find type of the value and go to newline
	fmt.Printf("%T \n","Hello")

	//example to round  value to 3 decimal and go to newline
	fmt.Printf("%.3f \n", 3.1329424)

	// example to print boolean value and go to newline
	fmt.Printf("%t \n",true)

	// example to print integer and go to newline
	fmt.Printf("%d \n", 12424)

	// example to print binary of a number
	fmt.Printf("%b \n", 9)

	// example to print ascii value for a given number
	fmt.Printf("%c \n", 65)
	fmt.Printf("%c \n", 97)

	// example to print hex value for given number
	fmt.Printf("%x \n", 15)
	fmt.Printf("%x \n", 16)
	fmt.Printf("%x \n", 42)
	fmt.Printf("%x \n", 160)

}
