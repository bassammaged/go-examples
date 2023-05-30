package dpinjection

import "fmt"

// Define the User interface
type User interface {
	AddUser(token string)
}

type NoPrivilegeUser struct{}

func (noPrivilegeUser NoPrivilegeUser) AddUser(token string) {
	if token == "noprivilege" {
		fmt.Println("Hello No Privilege")
	} else {
		fmt.Println("Doesn't matter, Welcome back!")
	}
}

type StandardUser struct{}

func (standardUser StandardUser) AddUser(token string) {
	if token == "standard" {
		fmt.Println("Hello Standard")
	} else {
		fmt.Println("You are faking the Standard privilege")
	}
}

type AdminUser struct{}

func (adminUser AdminUser) AddUser(token string) {
	if token == "admin" {
		fmt.Println("Hello Admin")
	} else {
		fmt.Println("You are faking the admin privilege")
	}
}

type AnonymousUser struct{}

func (anonUser AnonymousUser) AddUser(token string) {
	fmt.Println("Anonymous User :) Get out of here!")
}

// Define customer
type Customer struct {
	CustomerUser User
}

func NewCustomer(user User) *Customer {
	return &Customer{
		CustomerUser: user,
	}
}
