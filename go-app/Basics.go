package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var bookings []string

func main2() {
	var remainingTickets uint = 50
	greeting(remainingTickets)

	for {
		firstName, lastName, email, userTickets := userInput()
		isValidName, isValidEmail, isValidTicketsCounts := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketsCounts {
			bookTicket(firstName, lastName, email, remainingTickets, userTickets)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name and last name you entered are too short")
			}
			if !isValidEmail {
				fmt.Println("you have entered wrong email")
			}
			if !isValidTicketsCounts {
				fmt.Println("number of tickets entered is invalid")
			}
		}
	}
}

func greeting(remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Book your tickets here to attend")
}

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter no. of tickets wants to book: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsCounts := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketsCounts
}

func bookTicket(firstName string, lastName string, email string, remainingTickets uint, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Booking List:%v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Total tickets: %v, Remaining tickets: %v\n", conferenceTickets, remainingTickets)
}
