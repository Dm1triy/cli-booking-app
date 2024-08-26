package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func greetUser() {
	fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Printf("We have total of %v tickets and %v are still available.\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, string) {
	// uninitialized variable get default value
	// for the corresponding type
	var firstName string
	var lastName string
	var email string
	var inputTickets string
	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets you want to buy: ")
	fmt.Scan(&inputTickets)
	return firstName, lastName, email, inputTickets
}

func validateUserInput(firstName string, lastName string, email string, inputTickets string) (uint8, error) {

	var userTickets uint8

	i, inputError := strconv.Atoi(inputTickets)
	userTickets = uint8(i)

	isValidName := len(firstName) > 1 && len(lastName) > 1
	isValidEmail := strings.Contains(email, "@")
	isValidTiketNumber := userTickets > 0 && userTickets <= remainingTickets &&
		inputError == nil && i < int(^uint8(0))

	if !isValidName {
		fmt.Println("Entered first name or last name is invalid. Try book tickets again.")
		return 0, errors.New("invalid name")
	} else if !isValidEmail {
		fmt.Println("Entered email address is invalid. Try book tickets again.")
		return 0, errors.New("invalid email")
	} else if !isValidTiketNumber {
		fmt.Printf("You enetered ivalid number of tickets. We have only %v tickets available\n",
			remainingTickets)
		return 0, errors.New("invalid number of tickets")
	}
	return userTickets, nil
}

func bookTickets(firstName string, lastName string, email string, userTickets uint8) {
	/* for arrays

	bookings[0] = firstName + " " + lastName
	*/
	/* for slices

	bookings = append(bookings, firstName+" "+lastName)
	*/

	/* creates data structure with type map[key_type]val_type

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["ticketsNumber"] = strconv.FormatUint(uint64(userTickets), 10)
	*/

	var userData = userData{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		ticketsNumber: userTickets,
	}

	bookings = append(bookings, userData)
	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. "+
		"You will recieve a conformation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n",
		remainingTickets, conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var firstName = strings.Fields(booking)[0]
		// var firstName = booking["firstName"]
		var firstName = booking.firstName
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}
