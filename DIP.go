package main

import "fmt"

/*

//This code not follows Dependency Inversion Principle.
//What is DIP ? --> High Level Module should not depend on low level module. Instead they both depend on abstractions.
//what is high level module ? --> It provides us what it is going to do
//What s low level module? --> It provides us how we are doing

type User struct {
	userName      string
	MsgPreference string
}
type EmailNotifier struct{}

func (e *EmailNotifier) Notifier() {
	fmt.Println("Email Sent to User")
}

type SMSNotifier struct{}

func (s *SMSNotifier) Notifier() {
	fmt.Println("SMS sent to user")
}

func (u *User) NotifyUser() {           // Here NotifyUser() is the high level module
	switch u.MsgPreference {
	case "SMS":
		smsNotifier := &SMSNotifier{}   // SMSNotifier is low level module
		smsNotifier.Notifier()
	case "Email":
		emailNotifier := &EmailNotifier{} //EmailNotifier is low level module
		emailNotifier.Notifier()
	}
}

func main() {

	u1 := User{"User1", "SMS"}
	u1.NotifyUser()
	u2 := User{"User2", "Email"}
	u2.NotifyUser()
}

*/

type User struct {
	userName string
	notifier Notifier
}

type Notifier interface {
	Notify()
}

type EmailNotifier struct{}

func (e *EmailNotifier) Notify() {
	fmt.Println("Email Sent to User")
}

type SMSNotifier struct{}

func (s *SMSNotifier) Notify() {
	fmt.Println("SMS sent to user")
}

func (u *User) NotifyUser() {
	// So here the high level module NotifyUser() is depending on notifier interface
	// instead of directly depending onn smsnotifier and emailnotifier
	u.notifier.Notify()
}

func main() {

	u1 := User{"User1", &SMSNotifier{}}
	u1.NotifyUser()
	u2 := User{"User2", &EmailNotifier{}}
	u2.NotifyUser()
}
