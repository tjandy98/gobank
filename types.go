package main

import "math/rand"

type Account struct {
	ID uint `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Number uint `json:"number"`
	Balance uint `json:"balance"`
}

func NewAccount(firstName string, lastName string) *Account{
	return &Account{FirstName: firstName, LastName: lastName, ID: uint(rand.Intn(1000)), Number: uint(rand.Intn(1000000))}
}