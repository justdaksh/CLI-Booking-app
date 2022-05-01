package main

import (
	"fmt"
	"strings"
)
func main() {
	
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remaningTickets uint= 50

	fmt.Printf("Welcome to %v Booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are available.\n",conferenceTickets,remaningTickets)
	fmt.Println("Get your tickets here to attend")
	
	var firstName string
	var lastName string
	var emailId string
	var userTickets uint

	// array 
	var bookings []string
	
	//ask user for their username
	fmt.Println("Enter Your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter Your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scan(&emailId)

	fmt.Println("Enter no of tickets: ")
	fmt.Scan(&userTickets)

	remaningTickets -= userTickets
	bookings = append(bookings, firstName + " " + lastName)

	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v.\n",firstName,lastName,userTickets,emailId)
	
	firstNames := []string{}
	for _,booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("We have %v tickets left now, Hurry up!!\n",remaningTickets)
	fmt.Printf("FirstNames of booking are:  %v\n",firstNames)
}
