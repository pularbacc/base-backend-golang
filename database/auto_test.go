package database

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateAuto(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create an example Auto object
	auto := Auto{
		Name: "Example Auto 2",
		Cmds: []Cmd{
			{Method: "GET", Url: "/users", Body: ""},
			{Method: "GET", Url: "/history", Body: ""},
			{Method: "POST", Url: "/git/clone", Body: "{\"key\": \"value\"}"},
		},
	}

	// Create the Auto and Cmds in the database
	createdAuto, err := CreateAuto(db, auto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(createdAuto)
}

func TestGetAutoList(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	autos, err := GetAutoList(db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(autos)
}

func TestGetAutoById(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	auto, err := GetAutoByID(db, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(auto)
}

func TestUpdateAuto(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	auto := Auto{
		Id:   1,
		Name: "Example Auto Update",
		Cmds: []Cmd{
			{Method: "GET", Url: "/hehe", Body: ""},
		},
	}

	err = UpdateAuto(db, auto)

	if err != nil {
		log.Fatal(err)
	}
}

func TestDeleteAuto(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = DeleteAuto(db, 1)

	if err != nil {
		log.Fatal(err)
	}
}
