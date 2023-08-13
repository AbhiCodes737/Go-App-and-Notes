package printer

import "fmt"

func HeyThere(name string, totalTickets uint, remainingTickets uint){
	fmt.Println("\nWelcome to", name)
	fmt.Println("The Best Commercial Flight Service")
	fmt.Println("Total Tickets:", totalTickets)
	fmt.Println("Tickets Remaining:", remainingTickets)
}