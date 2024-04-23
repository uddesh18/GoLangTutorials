package main

import "fmt"

// Defining a Structure
type Address struct {
	name   string
	street string
	city   string
	state  string
	pin    int
}

// Nested Structures
type Person struct {
	pname     string
	plastname string
	age       int
	gender    string
	address   Address
}

func main() {

	var aStruct Address
	fmt.Println(aStruct)

	fmt.Println("Normal Structure")
	a := Address{"Uddesh", "Solapur", "Solapur", "MH", 12345}
	fmt.Println(a)

	fmt.Println("Nested Structue")
	person := Person{"Uddesh", "Takpere", 24, "M", a}
	fmt.Println(person.address.city)

	var conferenceName string = "Test Conference"
	const conferenceTickets int = 100
	var remainingTickets int = 100
	// Slice -> Abstraction of Array
	var bookings []string

	fmt.Printf("Name of conference is %v ", conferenceName)
	fmt.Println(" ")
	fmt.Printf("Total Number of Tickets are %v and remaining tickets are %v \n", conferenceTickets, remainingTickets)

	var userName string
	var userTickets int
	var email string

	// Reference Pointer - > passing memory address
	for {
		fmt.Scan(&userName)
		fmt.Scan(&userTickets)
		fmt.Scan(&email)

		bookings = append(bookings, userName)
		remainingTickets = remainingTickets - userTickets

		if remainingTickets <= 0 {
			for index := range bookings {
				fmt.Printf(bookings[index])
				break
			}
		}

		greetUser(userName)
		fmt.Printf("Username is %v , tickets required are %v , email is %v ", userName, userTickets, email)
		fmt.Printf("Remaining Tickets are : %v \n", remainingTickets)
		fmt.Printf("Slice type : %T \n ", bookings)
		fmt.Printf("Size of Slice : %v \n", len(bookings))
	}

}

func greetUser(userName string) {
	fmt.Printf("Hello %v , How are you ? \n", userName)
}
