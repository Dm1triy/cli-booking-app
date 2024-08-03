package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Package Level Variables - can be accessed inside any of the functions
const conferenceTickets uint8 = 50             // with this form (:=) const variables can't be defined
var conferenceName = "GO conference"           // as well as package level variables
var remainingTickets uint8 = conferenceTickets // as well as explicit type
var bookings []string                          // slice - dynamic array
/* array
var bookings [50]string
*/

func main() {
	// Welcome part
	greetUser()

	// Order part
	for remainingTickets > 0 {

		firstName, lastName, email, inputTickets := getUserInput()
		userTickets, InputError := validateUserInput(firstName, lastName, email, inputTickets)
		if InputError != nil {
			continue
		}

		bookTickets(firstName, lastName, email, userTickets)

		// instead of printing full name of the users who have booked
		// tikets, only their first names are printed
		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstNames)
	}

	fmt.Println("\nOur conference is booked out. Come back next year")

	learning_func := true
	if !learning_func {
		choose_city()
	}
}

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

	isValidName := len(firstName) > 1 && len(lastName) > 1 && inputError == nil && i < int(^uint8(0))
	isValidEmail := strings.Contains(email, "@")
	isValidTiketNumber := userTickets > 0 && userTickets <= remainingTickets

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
	// for slices
	bookings = append(bookings, firstName+" "+lastName)
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
		var firstName = strings.Fields(booking)[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func choose_city() {
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
