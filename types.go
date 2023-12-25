package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	Recipient int `json:"recipient"`
	Amount int `json:"amount"`

}
type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    uint      `json:"number"`
	Balance   uint      `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    uint(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}
