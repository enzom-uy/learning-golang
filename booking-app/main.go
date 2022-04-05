package main

import (
	"bufio"
	"fmt"
	"os"
)

var conferenceName = "Golang Conference" // Name of the conference
const conferenceTickets = 50             // Tickets available for the conference
var remainingTickets = conferenceTickets // Remaining tickets available for the conference
var ticketsBoughtByUser = 0              // Tickets bought by the user
var clientName string

func main() {
	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("Get your tickets here to attend.")
	fmt.Println("There are", remainingTickets, "tickets available.")
	fmt.Println("Please enter your name:")
	buyTickets()
}

func buyTickets() {
	// Read user input (name)
	inputName := bufio.NewReader(os.Stdin)
	userName, _ := inputName.ReadString('\n')
	clientName = userName // Save the name of the user

	// Creates the userTickets variable and sets it to 0
	var userTickets int
	fmt.Println("Hello,", clientName, "How many tickets do you want to buy?") // Ask user how many tickets they want to buy
	fmt.Scanln(&userTickets)                                                  // Read user input (number of tickets)

	// Asks the user to confirm the purchase
	fmt.Println("Do you confirm your purchase? (y/n)")
	var confirm string
	fmt.Scanln(&confirm)

	// If the user confirms the purchase, the tickets are bought and the remaining tickets are updated
	// If the user doesn't confirm the purchase, the tickets are not bought and exit the program
	if confirm == "y" {
		// Check if the are enough tickets available for the user,
		// if not, print an error message and exit the function, otherwise, continue.
		if userTickets > remainingTickets {
			fmt.Println("Sorry, we don't have enough tickets left!")
			return
		} else {
			fmt.Println("Thank you for your purchase! You have bought", userTickets, "tickets.")
			remainingTickets = remainingTickets - userTickets       // Subtract the number of tickets bought from the total available tickets
			ticketsBoughtByUser = ticketsBoughtByUser + userTickets // Add the number of tickets bought to the total tickets bought by the user
		}
	} else {
		fmt.Println("Thank you for your interest in", conferenceName, "!")
		return
	}
	fmt.Println("You have bought", ticketsBoughtByUser, "tickets")
	fmt.Println("There are now", remainingTickets, "tickets left")

}
