package database

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	db, err := Connect()

	if err != nil {
		log.Fatal(err)
	}

	profile, err := CreateProfile(db, ProfileReq{
		Name:       "Cong Nguyen",
		Account_id: 6,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profile)
}
