package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go conference"
	const confernaceTicket = 50
	var remainingTicket uint = 50
	fmt.Printf("Hello..Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v still available\n", confernaceTicket, remainingTicket)
	fmt.Println("Get you tickets from here to attend")
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//array
	//var bookings [50]string
	//sslice
	bookings := []string{}
	for {
		fmt.Println("Enter you first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter you Last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter you email")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets you want")
		fmt.Scan(&userTickets)
		//example of array
		//bookings[0] = firstName + " " + lastName
		//example of slice
		if userTickets > remainingTicket {
			fmt.Printf("We have only %v tickets remaining,So you can't book %v tikets", remainingTicket, userTickets)
			break
		}
		remainingTicket = remainingTicket - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		//fmt.Printf("The whole slice: %v\n", bookings)
		//fmt.Printf("The whole slice: %v\n", bookings[0])
		//fmt.Printf("  Type slice: %T\n", bookings)
		//fmt.Printf("The  arraylength: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v of tickets.You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v remaining out of%v\n", remainingTicket, confernaceTicket)
		firstNames := []string{}
		//for index, booking := range bookings {
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("The firstnames of  bookings are : %v\n", firstNames)
		var noTicketsRemaining bool = remainingTicket == 0
		if noTicketsRemaining {
			//end program
			fmt.Println("Our conference is booked.please come next year")
			break

		}
	}
}
