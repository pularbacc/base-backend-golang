package database

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateToken(t *testing.T) {
	db, err := Connect()

	if err != nil {
		log.Fatal(err)
	}

	idToken, err := CreateToken(db, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(idToken)
}

func TestCreateRefreshToken(t *testing.T) {
	db, err := Connect()

	if err != nil {
		log.Fatal(err)
	}

	idToken, err := CreateRefreshToken(db, Token{
		Account_id: 1,
		Token_id:   1,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(idToken)
}

func TestGetAccountByToken(t *testing.T) {
	db, err := Connect()

	if err != nil {
		log.Fatal(err)
	}

	acc, err := GetAccountByToken(db, 3)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(acc)
}

func TestGetInfoRefreshToken(t *testing.T) {
	db, err := Connect()

	if err != nil {
		log.Fatal(err)
	}

	info, err := GetInfoRefreshToken(db, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
}
