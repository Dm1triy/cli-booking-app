package choosecity

import "fmt"

// In order to export function or variable and make it
// available for all packages in the app, the first
// letter needs to be capitalized
func Choose_city() {
	var city string
	fmt.Scan(&city)
	switch city {
	case "New York":
		fmt.Println("Your country is USA")
	case "London":
		fmt.Println("Your country is England")
	case "Berlin", "Munich":
		fmt.Println("Your country is Germany")
	default:
		fmt.Println("Your city not listed here")
	}
}
