package printer

import "fmt"

func ConditionChecker(isValidName bool, isValidEmail bool, isValidTicket bool, remainingTickets uint){
	if !isValidName {
		fmt.Println("First Name or Last Name must be at least two characters long. Try again")
	}
	if !isValidEmail {
		fmt.Println("Invalid Email Address. Try again")
	}
	if !isValidTicket {
		fmt.Printf("Only %v ticket(s) remaining. Try again\n", remainingTickets)
	}
}