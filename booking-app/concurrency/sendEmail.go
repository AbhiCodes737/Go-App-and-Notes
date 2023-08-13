package concurrency

import (
	"booking-app/structData"
	"fmt"
	"sync"
	"time"
)

func SendEmail(deets structData.UserData, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	var message = fmt.Sprintf("CONFIRMED: %v ticket(s) for %v %v\n", deets.Tickets, deets.FirstName, deets.LastName)
	fmt.Println("--------------------------------------------")
	fmt.Println("FROM: mail@flygo.com")
	fmt.Printf("TO: %v\n", deets.Email)
	fmt.Println("WELCOME ONBOARD!!!")
	fmt.Print(message)
	fmt.Println("--------------------------------------------")
	wg.Done()
}
