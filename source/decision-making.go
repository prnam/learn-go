package main

import "fmt"

func main()  {
		statusCheck(200)
		statusCheck(404)
		statusCheck(301)
		statusCheck(106)
		statusCheck(500)
		statusCheck(600)

}

func statusCheck(apiResponseStatus int)  {
	fmt.Println("--- Executing statusCheck ---")
	if apiResponseStatus >= 100 && apiResponseStatus < 200{
		fmt.Println("Informational responses")
	} else if apiResponseStatus >= 200 && apiResponseStatus < 300{
		fmt.Println("Successful responses")
	} else if apiResponseStatus == 301 && apiResponseStatus < 400{
		fmt.Println("Redirects")
	} else if apiResponseStatus >= 400 && apiResponseStatus < 500{
		fmt.Println("Client side errors")
	}  else if apiResponseStatus == 500 {
		fmt.Println("Server side errors")
	} else {
		fmt.Println("Seriously?")
	}
	if apiResponseStatus < 600 { checkSelectedStatus(apiResponseStatus) }
}

func checkSelectedStatus(apiResponseStatus int){
	fmt.Println("--- Executing checkSelectedStatus ---")
	message :=""
	switch apiResponseStatus{
	case 100: message = "Continue"
	case 200: message = "OK"
	case 301: message = "Moved Permanently"
	case 404: message = "Not found"
	case 500: message = "Internal Server Error"
	default: message = "This example is for learning switch syntax in GO, so limited to 5 common http code." +
		"Try only 100, 200, 301, 404 and 500 HTTP code"
	}
	fmt.Println(message)
}