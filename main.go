package main

import (
	"fmt"
	"strconv"
	"strings"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remaningTickets uint = 50
var bookings = make([]map[string]string,0) 

func main() {

	greetUsers()
	
	
	for {
		//ask user for their username
		
		firstName, lastName, emailId, userTickets := getUserInput()
		
		isValidEmail, isValidName, isValidTickets := validateUserInput(firstName,lastName,emailId,userTickets)

		if isValidEmail && isValidName && isValidTickets {

			bookTicket(userTickets,firstName,lastName,emailId)
			
			firstNames := getFirstNames()
			
			fmt.Printf("FirstNames of booking are:  %v\n",firstNames)
	
			noTickets := remaningTickets == 0
			if noTickets {
				fmt.Println("Our Conference is booked out. Come back next year!!")
				break
			}
		}else { 
			if !isValidEmail {
				fmt.Println("Please Enter a valid Email")
			}
			if !isValidName {
				fmt.Println("Please Enter a valid Name")
			}
			if !isValidTickets {
				fmt.Printf("Sorry Only %v Tickets are available.\n", remaningTickets)
			}
		}
		
	}
}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are available.\n",conferenceTickets,remaningTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _,booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func validateUserInput(firstName string,lastName string,emailId string,userTickets uint) (bool,bool,bool) {

	isValidName := len(firstName)>=1 && len(lastName)>=1

	isValidEmail := strings.Contains(emailId , "@") && strings.Contains(emailId, ".com")

	isValidTickets := userTickets > 0 && userTickets <= remaningTickets
	
	return isValidEmail , isValidName , isValidTickets

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailId string
	var userTickets uint

	fmt.Println("Enter Your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter Your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scan(&emailId)

	fmt.Println("Enter no of tickets: ")
	fmt.Scan(&userTickets)

	return firstName,lastName,emailId,userTickets
}

func bookTicket(userTickets uint,firstName string,lastName string,emailId string) {

	remaningTickets -= userTickets
	var userData =make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = emailId
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10)
	
	bookings = append(bookings,userData)		
	fmt.Printf("List of bookings is %v\n",userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v.\n",firstName,lastName,userTickets,emailId)
	fmt.Printf("We have %v tickets left now for %v, Hurry up!!\n",remaningTickets,conferenceName)
}