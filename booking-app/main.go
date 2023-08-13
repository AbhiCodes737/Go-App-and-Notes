package main

import (
	"booking-app/concurrency"
	"booking-app/printer"
	"booking-app/structData"
	"fmt"
	"strings"
	"sync"
)

// Concurrency Synchronisation
var wg = sync.WaitGroup{}

// Package Variables
var name = "Flygo"

const totalTickets uint = 100

var remainingTickets uint = 100
var bookings = make([]structData.UserData, 0)

func main() {

	printer.HeyThere(name, totalTickets, remainingTickets)

	// Local Variables
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for remainingTickets > 0 {

		//User Input
		fmt.Print("\nEnter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// User Validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		// Condition Checking
		if isValidName && isValidTicket && isValidEmail {

			// Create a struct for the user
			var userDetails = structData.UserData{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Tickets:   userTickets,
			}

			// Calculation
			bookings = append(bookings, userDetails)
			fmt.Printf("\nThank you %v %v for booking %v tickets. The confirmation link will be mailed to %v.\n\n", firstName, lastName, userTickets, email)
			remainingTickets -= userTickets

			// Concurrent Email Send with reference to waitgroup
			wg.Add(1) // Adds thread to wg
			go concurrency.SendEmail(userDetails, &wg)

			// Function Output
			printer.GetFirstNames(bookings)

			// Final ticket check and end loop
			if remainingTickets == 0 {
				fmt.Println("\nWe are fully booked")
				break
			}
		} else {
			// Validate Input
			printer.ConditionChecker(isValidName, isValidEmail, isValidTicket, remainingTickets)
			continue
		}
	}
	wg.Wait() // Waits for coroutine
}
