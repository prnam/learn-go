package main

import "fmt"

func main(){
	fmt.Println("main: Starts at main.. ⏭")
	someFunction()
	fmt.Println("main: OK, everybody calm down 🧘️")
}

func someFunction() {
	defer func() {
		if result := recover(); result != nil {
			fmt.Println("someFunction: Flow recovered to normal, i assume someone panicked 🤔", result)
		}
	}()
	var val bool = false
	fmt.Println("someFunction: Executing function called someOtherFunction to simulate non panic condition...⌛️")
	someOtherFunction(val)
	fmt.Println("someFunction: Returned normal after calling someOtherFunction 👍")
	val = true
	fmt.Println("someFunction: calling someOtherFunction to simulate panic condition.. 🤙")
	someOtherFunction(val)
	fmt.Println("someFunction: Flow looks normal after calling someOtherFunction 👍")
}

func someOtherFunction(val bool) {
	if val == true{
		panic(fmt.Sprintf("\nsomeOtherFunction: I panicked 😱"))
	} else {
		fmt.Println("someOtherFunction: Executing else block business logic ⌛️")
	}
}




