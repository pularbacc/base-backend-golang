package database

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	id, err := CreateAccount(db, AccountCreateReq{
		Email:    "abc@gmail.com",
		Password: "123",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}

func TestGetListAccount(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	accounts, err := GetListAccount(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(accounts)
}

func TestGetAccount(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	account, err := GetAccountInfoById(db, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account)
}

func TestGetAccountByEmail(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	account, err := GetAccountInfoByEmail(db, "abc@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account)
}

func TestGetPassword(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	pass, err := GetPassByEmail(db, "abc@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pass)
}
