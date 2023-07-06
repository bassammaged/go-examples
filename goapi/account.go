package main

import "math/rand"

type Account struct {
	Id             int    `json:"id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Age            uint8  `json:"age"`
	HashedPassword string `json:"hashedPassword"`
	Email          string `json:"email"`
	ApiKey         string `json:"apiKey"`
}

type userSession struct {
	loggedIn bool
	jwtToken string
}

func NewAccount() *Account {
	return &Account{
		Id:        rand.Intn(10000),
		Firstname: "Bassam",
		Lastname:  "Kemet",
		Email:     "k.b@gmail.com",
	}
}

func (u *Account) CreateAccount(firstName, lastName, email, password string, age uint8) error {
	// Do some checks
	if err := u.doChecks(); err != nil {
		return err
	}
	// Set the new values
	// Store user data

	return nil
}

func (u *Account) GetAccount(id int) error {
	// Do some checks
	if err := u.doChecks(); err != nil {
		return err
	}

	// get the account information
	u.Id = id
	u.Firstname = "Bassam"
	u.Lastname = "Kemet"
	u.Email = "b@hotmail.com"

	return nil
}

func (u Account) doChecks() error {
	return nil
}
