package printer

import (
	"booking-app/structData"
	"fmt"
)

func GetFirstNames(bookings []structData.UserData) {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}
	fmt.Print("Users booked till now: ")
	fmt.Println(firstNames)

}
