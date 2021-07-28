package main

import "fmt"

func main()  {
	valueOne, valueTwo := 22, 7
	fmt.Println("--- Arithmetic Operators Example ---")
	fmt.Println("Add",valueTwo,"to",valueOne,":", valueOne+valueTwo)
	fmt.Println("Subtract",valueTwo,"from",valueOne,":", valueOne-valueTwo)
	fmt.Println("Multiply",valueOne,",",valueTwo,"times :", valueOne*valueTwo)
	fmt.Println("Quotient of",valueOne,"divide by",valueTwo,":",valueOne/valueTwo)
	fmt.Println("Remainder of",valueOne,"divide by",valueTwo,":",valueOne%valueTwo)

	fmt.Println("--- Relational Operators Example ---")
	fmt.Println(valueOne,"greater than",valueTwo,": ", valueOne > valueTwo)
	fmt.Println(valueOne,"greater than or equal to",valueTwo,": ", valueOne >= valueTwo)
	fmt.Println(valueOne,"lesser than",valueTwo,": ", valueOne < valueTwo)
	fmt.Println(valueOne,"lesser than or equal to",valueTwo,": ", valueOne <= valueTwo)
	fmt.Println(valueOne,"equal to",valueTwo,": ", valueOne == valueTwo)
	fmt.Println(valueOne,"not equal to",valueTwo,": ", valueOne != valueTwo)

	_true := true
	_false := false
	fmt.Println("--- Logical Operators Example ---")
	fmt.Println(_true,"and",_false,": ", _true && _false)
	fmt.Println(_true,"or",_false,": ", _true || _false)
	fmt.Println("Negate of", _false,": ", !_false)
	fmt.Println("Negate of", _true,": ", !_true)

}
