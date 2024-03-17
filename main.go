package main

import (
	"booking-app/helper" // Importing helper package for input validation
	"fmt"
	"sync"
	"time"
	// "strconv"
)

// Total tickets available for conference
const conferenceTicket int = 50

// Name of the conference
var conferenceName string = "Go Conference"

// Number of remaining tickets
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)

var bookings = make([]UserData, 0)

// Struct to hold user booking data
type UserData struct {
	userName        string // User's name
	email           string // User's email
	numberOfTickets uint   // Number of tickets booked by the user
}

// Function to greet user and display available tickets
func greetUser() {
	fmt.Println("Welcom to", conferenceName, "booking application")
	fmt.Println("------------------------------------------------------")
	fmt.Printf("We have total of %v ticket and %v ticket are available\n", conferenceTicket, remainingTickets)
}

// Function to take user input for booking
func usetInput() (string, string, uint) {
	var userName string
	var email string
	var userTicket uint
	fmt.Println("Please enter your name")
	fmt.Scan(&userName)

	fmt.Println("Please enter your email")
	fmt.Scan(&email)

	fmt.Println("Please enter how many ticket to book")
	fmt.Scan(&userTicket)
	return userName, email, userTicket
}

// Function to book tickets and update remaining tickets
func bookTicket(userName string, userTicket uint, email string) {
	remainingTickets -= userTicket

	//Create a map for user booking

	// var userData = make(map[string]string)
	// userData["userName"] = userName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)

	//Create a Struct
	var userData = UserData{
		userName:        userName,
		email:           email,
		numberOfTickets: userTicket,
	}

	// Append user booking data to the bookings slice
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v to book %v ticket. You will receive a confirmation email at %v\n", userName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// WaitGroup to wait for all goroutines to finish
var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for { // for is infinite loop and is equal  for remainingTickets < 50 is same

		userName, email, userTicket := usetInput()

		isValiedName, isValidEmail, isValiedUserTicket := helper.ValidateUserInput(userName, email, userTicket, remainingTickets)

		if isValiedName && isValidEmail && isValiedUserTicket {

			bookTicket(userName, userTicket, email)

			wg.Add(1)
			go sendTicket(userName, userTicket, email)

			for _, booking := range bookings {
				// fmt.Printf("These are the booking %s\n", booking["userName"])
				fmt.Printf("These are the booking %s\n", booking.userName)
			}
			if remainingTickets == 0 {
				fmt.Println("Our conference is Full Booked!")
				break
			}
		} else {
			// only if because the user might be enterd all 3 wrong
			if !isValiedName {
				fmt.Println("Please enter valid name")
			}
			if !isValidEmail {
				fmt.Println("Please enter valid email")
			}
			if !isValiedUserTicket {
				fmt.Printf("We only have %v ticket ramining\n", remainingTickets)
			}
		}
		fmt.Println("---------------------------------")

	}
	wg.Wait()
}

// Function to simulate sending ticket (dummy function)
func sendTicket(userName string, userTicket uint, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v ", userTicket, userName)
	fmt.Println("*****************************************************")
	fmt.Printf("Sending ticket: \n %v \n to email %v \n", ticket, email)
	fmt.Println("*****************************************************")
	wg.Done()
}
