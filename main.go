package main

import (
	"fmt"
	"strings"
)

const confernaceTicket = 50

var conferenceName = "Go conference"
var remainingTicket uint = 50

// sslice
var bookings = []string{}

func main() {

	greetUser()

	//array
	//var bookings [50]string

	for {
		firstName, lastName, email, userTickets := getUserInput()
		//example of array
		//bookings[0] = firstName + " " + lastName
		isValidName, isValidEmail, isValidTicketsNumber := validateUser(firstName, lastName, email, userTickets)
		//example of slice
		if isValidName && isValidEmail && isValidTicketsNumber {
			//call bookticketfunc
			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The whole slice: %v\n", bookings[0])
			//fmt.Printf("  Type slice: %T\n", bookings)
			//fmt.Printf("The  arraylength: %v\n", len(bookings))
			bookTicket(remainingTicket, userTickets, bookings, firstName, lastName, email, conferenceName)
			//for index, booking := range bookings {
			//call print first name
			firstNames := getFirstName()
			fmt.Printf("The firstnames of  bookings are : %v\n", firstNames)
			var noTicketsRemaining bool = remainingTicket == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is booked.please come next year")
				break

			}

		} else {
			if !isValidName {
				fmt.Println("first name and last name is to short")
			}
			if !isValidEmail {
				fmt.Println("fNot conatain @ symbol")
			}
			if !isValidTicketsNumber {
				fmt.Println("Number of tickets is invalid")
			}
			//fmt.Println("Your input data is invalid")
			//continue
		}
	}
}

func greetUser() {
	fmt.Printf("Hello..Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v still available\n", confernaceTicket, remainingTicket)
	fmt.Println("Get you tickets from here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}
func validateUser(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTicket
	return isValidName, isValidEmail, isValidTicketsNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter you first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter you Last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter you email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}
func bookTicket(remainingTicket uint, userTickets uint, bookings []string, firstName string, lastName string, email string, confernaceTicket string) {
	remainingTicket = remainingTicket - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v of tickets.You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining out of%v\n", remainingTicket, confernaceTicket)
}
