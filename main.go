package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// with this form (:=) const variable can't be defined
	// as well as explicit type
	conferenceName := "GO conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = conferenceTickets
	/* array

	var bookings [50]string

	*/
	// slice - dynamic array
	var bookings []string

	// Welcome part
	fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Printf("We have total of %v tickets and %v are still available.\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Order part
	for remainingTickets > 0 {
		// uninitialized variable get default value
		// for the corresponding type
		var firstName string
		var lastName string
		var email string
		var inputTickets string
		var userTickets uint8
		fmt.Print("\nEnter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Print("Enter number of tickets you want to buy: ")
		fmt.Scan(&inputTickets)

		i, inputError := strconv.Atoi(inputTickets)
		userTickets = uint8(i)

		isValidName := len(firstName) > 1 && len(lastName) > 1 &&
			inputError == nil && i < int(^uint8(0))
		isValidEmail := strings.Contains(email, "@")
		isValidTiketNumber := userTickets > 0 && userTickets <= remainingTickets

		if !isValidName || !isValidEmail || !isValidTiketNumber {
			fmt.Println("Your input data is invalid, try again.")
			continue
		}

		/* for arrays

		bookings[0] = firstName + " " + lastName

		*/
		// for slices
		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets -= userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. "+
			"You will recieve a conformation email at %v\n",
			firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n",
			remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var firstName = strings.Fields(booking)[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)
	}

	fmt.Println("\nOur conference is booked out. Come back next year")
}
