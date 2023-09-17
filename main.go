package main

import "fmt"

func main() {

	var conferenceName string = "Go conference"
	const confernaceTicket = 50
	var remainingTicket uint = 50
	fmt.Printf("Hello..Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v still available\n", confernaceTicket, remainingTicket)
	fmt.Println("Get you tickets from here to attend")
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings [50]string
	fmt.Println("Enter you first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter you Last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter you email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want")
	fmt.Scan(&userTickets)
	bookings[0] = firstName + " " + lastName
	remainingTicket = remainingTicket - userTickets
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The whole array: %v\n", bookings[0])
	fmt.Printf("  Type array: %T\n", bookings)
	fmt.Printf("The  arraylength: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v of tickets.You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining out of%v\n", remainingTicket, confernaceTicket)
}
