package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceName string = "Go Conference";
const conferenceTickets int = 50;
var remainingTickets uint = 50;
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers();

	firstName, lastName, email, userTickets := getUserInput();

	isValidName, isValidEmail, isValidTicketNumber, isValidBooking := helper.ValidateRegistrationForm(firstName, lastName, email, userTickets, remainingTickets);

	if isValidBooking {
		bookTicket(userTickets, firstName, lastName, email);

		wg.Add(1);
		go sendTicket(userTickets, firstName, lastName, email);

		firstNames := getFirstNames();
		fmt.Printf("These are all our bookings: %v \n", firstNames);
		
		var noTicketsRemaining bool = remainingTickets == 0;

		if noTicketsRemaining {
			fmt.Println("Our conference is booked out! Sorry! We'll see you in 2023");
			// break;
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short");
		}

		if !isValidEmail {
			fmt.Println("Email address you entered does not contain @ sign");
		}

		if !isValidTicketNumber {
			fmt.Println("The number of tickets you entered is invalid");
		}

		fmt.Println("Your input data is invalid, try again");
	}

	wg.Wait();
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName);
	fmt.Printf("We have a total of %v tickets and %v tickets are still remaining!\n", conferenceTickets, remainingTickets);
	fmt.Println("Get your tickets to attend!");
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets;
}

func getFirstNames() []string {
	firstNames := []string{}
	
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName);
	}

	return firstNames;
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets;

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData);
	fmt.Printf("List of bookings is %v\n", bookings);

	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v \n", firstName, lastName, userTickets, email);
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName);
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second);
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName);
	fmt.Println("###############");
	fmt.Printf("Sending ticket %v to email address %v", ticket, email);
	fmt.Println("###############");
	wg.Done();
}