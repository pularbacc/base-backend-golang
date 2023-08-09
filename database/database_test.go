package database

import (
	"fmt"
	"log"
	"testing"
)

func TestDatabase(t *testing.T) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
}
