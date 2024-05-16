package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName = "Go Conference"
	// conferenceName := "Go Conference"
	// const conferenceTicket = 50
	// var remainingTickets = 50
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Println("We have total of", conferenceTicket, "tickets and", remainingTickets, "are still available.")
	// fmt.Println("Get your tickets here to attend")

	// // var userName string
	// // var userTickets int
	// // // fmt.Printf("userName type %T, userTickets type %T \n", userName, userTickets)
	// // // userName = "Tom"
	// // // userTickets = 2

	// // // ask user for their name
	// // fmt.Scan(&userName) // use pointer

	// // // check pointer
	// // // fmt.Println(remainingTickets)	// 50
	// // // fmt.Println(&remainingTickets)	// 0x14000112018

	// // fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	// var remainingTickets int = 50
	// var firstName string
	// var lastName string
	// var email string
	// var userTickets int
	// // ask user info
	// fmt.Println("Enter your first name: ")
	// fmt.Scan(&firstName)

	// fmt.Println("Enter your last name: ")
	// fmt.Scan(&lastName)

	// fmt.Println("Enter your email name: ")
	// fmt.Scan(&email)

	// fmt.Println("Enter your ticket count: ")
	// fmt.Scan(&userTickets)

	// fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation email at", email)
	// // Enter your first name:
	// // jh
	// // Enter your last name:
	// // sim
	// // Enter your email name:
	// // nn@nn.nn
	// // Enter your ticket count:
	// // 5
	// // Thank you jh sim for bokking 5 tickets. You will receive a confirmation email at nn@nn.nn
	// // var bookings [50]string
	// var bookings []string // list

	// // bookings[0] = firstName + " " + lastName
	// bookings = append(bookings, firstName+" "+lastName)
	// remainingTickets = remainingTickets - userTickets
	// fmt.Println("RemainingTickets: ", remainingTickets)
	// fmt.Println("Bookings: ", bookings)

	// loop
	var bookings []string // list
	var remainingTickets int = 50
	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		// ask user info
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email name: ")
		fmt.Scan(&email)

		fmt.Println("Enter your ticket count: ")
		fmt.Scan(&userTickets)

		fmt.Println("Thank you", firstName, lastName, "for booking", userTickets, "tickets. You will receive a confirmation email at", email)

		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - userTickets
		fmt.Println("RemainingTickets: ", remainingTickets)
		// fmt.Println("Bookings: ", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Println("First Names: ", firstNames)

		// if remainingTickets <= 0 {
		noTicketsRemaining := remainingTickets <= 0
		if noTicketsRemaining {
			// end program
			fmt.Println("Out conference is booked out. Come back next year.")
			break
		}
	}
}
