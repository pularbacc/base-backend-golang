package database

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateRole(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	c, err := CreateRole(db, "ROOT")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}

func TestMountRole(t *testing.T) {
	db, err := Connect()

	if err != nil {
		log.Fatal(err)
	}

	id, err := MountRole(db, 6, 2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)

}
