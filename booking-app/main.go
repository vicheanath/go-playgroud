package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // unsigned integer
	var bookings[50] string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	



	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ast user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n" ,bookings)
	fmt.Printf("The first value: %v\n" ,bookings[0])
	fmt.Printf("The first value: %T\n" ,bookings)
	fmt.Printf("The array length: %v\n" ,len(bookings))


	fmt.Printf("Thanks you %v %v for booking %v tickets. You will receive a confirmation email at %v \n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets,conferenceName)
	
}
