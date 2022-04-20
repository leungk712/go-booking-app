package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName string = "Go Conference";
	const conferenceTickets int = 50;
	var remainingTickets uint = 50;
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName);
	fmt.Printf("We currently have %v tickets remaining!\n", remainingTickets);
	fmt.Println("Get your tickets to attend!");


	for {
		var firstName string;
		var lastName string;
		var email string;
		var userTickets uint;
	
		fmt.Println("Please enter your first name:");
		fmt.Scan(&firstName);
	
		fmt.Println("Please enter your last name:");
		fmt.Scan(&lastName);
	
		fmt.Println("Please enter your email:");
		fmt.Scan(&email);
	
		fmt.Println("Please enter how many tickets:");
		fmt.Scan(&userTickets);
	
		remainingTickets = remainingTickets - userTickets;
		bookings = append(bookings, firstName + " " + lastName);
	
		fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v \n", firstName, lastName, userTickets, email);
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName);
	
		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking);
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all our bookings: %v \n", firstNames);
	}
}
