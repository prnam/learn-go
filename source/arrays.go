package main

import "fmt"

func main() {
	regionsInIndia := []string{"Central","North","South","East","West"}
	forLoopUsingIndexValue(regionsInIndia)

	regionsInIndiaWithAvailabilityZones := make([]string,2)
	copy(regionsInIndiaWithAvailabilityZones, regionsInIndia)
	forLoopUsingIndexValue(regionsInIndiaWithAvailabilityZones)

	regionsInIndia = append(regionsInIndia, "North-East","South-West")
	forLoopUsingIndexValue(regionsInIndia)
}


func forLoopUsingIndexValue(regions []string)  {
	fmt.Println(" -- Executing for loop using range --")
	fmt.Println(regions)
	for index, value:= range regions{
		fmt.Println(index, value)
	}
}