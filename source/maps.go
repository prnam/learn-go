package main

import (
	"fmt"
)

func main() {
	organisation := make(map[int] string)
	organisation[1001] = "GitHub"
	organisation[1013] = "LinkedIn"

	fmt.Println(organisation)
	fmt.Println(len(organisation))

	microsoft := map[string]map[string]string{
		"GitHub": map[string]string{
			"ID": "ORG1001",
			"City":   "SC",
			"Target": "Developers",
		},
		"LinkedIn": map[string]string{
			"ID": "ORG1031",
			"City":   "SC",
			"Target": "Professionals",
		},
	}
	fmt.Println(microsoft)
	fmt.Println(len(microsoft))

	fmt.Println(microsoft["GitHub"])
	fmt.Println(len(microsoft["GitHub"]))

	fmt.Println(microsoft["LinkedIn"])
	fmt.Println(len(microsoft["LinkedIn"]))

	if item , businessUnit := microsoft["GitHub"]; businessUnit{
		fmt.Println(item["City"],item["Target"])
	}
}
