package main

import "fmt"

type UserDetails struct {
	Name        string
	Email       string
	PhoneNumber string
}

var userDetails []UserDetails

func SaveUser(u UserDetails) {
	userDetails = append(userDetails, u)
	fmt.Printf("%v details  saved \n", u.Name)
}

func SendMail(u UserDetails) {
	fmt.Printf("Mail sent to %v  for %v \n", u.Email, u.Name)
}

func main() {

	/* The below commented code doesn't follow SRP as main is function is responsible for creating,saving user and sending mail.


	//user1 := UserDetails{"User1", "user1@gmail.com", "9876543210"}
	//
	//userDetails = append(userDetails, user1)
	//fmt.Printf("Details  of %v are saved \n", user1.Name)
	//fmt.Printf("Success Email sent to %v \n", user1.Name)
	//
	//user2 := UserDetails{"User2", "user2@gmail.com", "98732099708"}
	//userDetails = append(userDetails, user2)
	//fmt.Printf("Details  of %v are saved \n", user2.Name)
	//fmt.Printf("Success Email sent to %v \n", user2.Name)

	*/

	//The below follows SRP we are just making calls not implementing the logic , main function responsibility is to call the functions , so it doesn't make main in doing multiple tasks

	user1 := UserDetails{"User1", "user1@gmail.com", "9876543210"}
	SaveUser(user1)
	SendMail(user1)

	user2 := UserDetails{"User2", "user2@gmail.com", "98732099708"}
	SaveUser(user2)
	SendMail(user2)

}
