package model

type User struct {
	Id int
	Email string
	EncryptedPassword string
	AccountRub int
	AccountUsd int
	AccountEur int
	AccountBalanceRub int
	AccountBalanceUsd int
	AccountBalanceEur int
}
