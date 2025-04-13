package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstname       string
	lastname        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstname, lastname, email, userTickets := userInput()
	isValidName, isValidEmail, isValidTickets := helper.UserInputValidation(firstname, lastname, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {
		remainingTickets = remainingTickets - userTickets

		userData := UserData{
			firstname:       firstname,
			lastname:        lastname,
			email:           email,
			numberOfTickets: userTickets,
		}

		bookings = append(bookings, userData)
		wg.Add(1)
		go sendTicket(userTickets, firstname, lastname, email)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstname, lastname, userTickets, email)
		fmt.Printf("%v tickets are remaining\n", remainingTickets)

		firstNames := getFirstName()
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		fmt.Printf("Entered Data: %v\n", userData)
		fmt.Printf("List of Bookings: %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Println("Tickets are Sold Out")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("firstname or lastname you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTickets {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("conferenceName: %T, conferenceTickets %T\n", conferenceName, conferenceTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and are %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstname)
	}
	return firstNames
}

func userInput() (string, string, string, uint) {
	var firstname, lastname, email string
	var userTickets uint

	fmt.Println("Enter the firstname:")
	fmt.Scan(&firstname)

	fmt.Println("Enter the lastname:")
	fmt.Scan(&lastname)

	fmt.Println("Enter the email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func sendTicket(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v ticket for %v %v", userTickets, firstname, lastname)
	fmt.Println("###############################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("###############################")
	wg.Done()
}
