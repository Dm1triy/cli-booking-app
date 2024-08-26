package main

import (
	choosecity "booking_app/choose_city"
	"fmt"
	"sync"
)

// Package Level Variables - can be accessed inside any of the functions
const conferenceTickets uint8 = 50             // with this form (:=) const variables can't be defined
var conferenceName = "GO conference"           // as well as package level variables
var remainingTickets uint8 = conferenceTickets // as well as explicit type
var bookings = make([]userData, 0)             // slice of custom type

/* slice - dynamic array. It's creating a slice of maps
   with zero initial size

var bookings = make([]map[string]string, 0)
*/

/* array

var bookings [50]string
*/

// struct allows to use fields with mixed type
// to define new type keyword "type" is used
type userData struct {
	firstName     string
	lastName      string
	email         string
	ticketsNumber uint8
}

// without WaitGroup when program is finished and
// some goroutines is not, then these routines will be aborted
var wg = sync.WaitGroup{}

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

		// Add method sets the number of goroutines to wait for
		wg.Add(1)
		// go keyword make function concurrent -
		// additional goroutine (lightweight thread) is created for that function
		go sendTicket(firstName, lastName, email, userTickets)

		// instead of printing full name of the users who have booked
		// tikets, only their first names are printed
		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstNames)
	}

	fmt.Println("\nOur conference is booked out. Come back next year")

	learning_func := true
	if !learning_func {
		choosecity.Choose_city()
	}
	// blocks main thread until the WaitGroup couter is 0
	wg.Wait()
}
