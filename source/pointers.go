package main

import "fmt"

func main()  {
	variable :="New"
	fmt.Println("--- executing main function ---")
	fmt.Println("Value of variable in main function ->",variable)
	fmt.Println("Address of variable in main function->",&variable)

	another_function(variable)
	fmt.Println("--- executing main function ---")
	fmt.Println("Value of variable in main function after executing another_function ->",variable)
	fmt.Println("Address of variable in main function after executing another_function ->",&variable)

	yet_another_function(&variable)
	fmt.Println("--- executing main function ---")
	fmt.Println("Value of variable in main function after executing yet_another_function ->",variable)
	fmt.Println("Address of variable in main function after executing yet_another_function ->",&variable)
}

func another_function(variable string)  {
	fmt.Println("--- executing another_function ---")
	fmt.Println("Value of variable in another_function before reassigning value ->",variable)
	fmt.Println("Address of variable in another_function before reassigning value ->",&variable)
	variable = "Updated"
	fmt.Println("Value of a variable in another_function after reassigning value ->",variable)
	fmt.Println("Address of variable in another_function after reassigning value ->",&variable)
}

func yet_another_function(variable *string)  {
	fmt.Println("--- executing yet_another_function ---")
	fmt.Println("Value of variable in yet_another_function before reassigning value ->",variable)
	fmt.Println("Address of variable in yet_another_function before reassigning value ->",&variable)
	*variable = "Updated"
	fmt.Println("Value of a variable in yet_another_function after reassigning value ->",variable)
	fmt.Println("Address of variable in yet_another_function after reassigning value ->",&variable)
}
