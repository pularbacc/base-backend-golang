package database

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/go-sql-driver/mysql"
)

func TestAlbum(t *testing.T) {
	var db *sql.DB

	// Capture connection properties.
	cfg := mysql.Config{
		User:   "pular",
		Passwd: "pular123",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "recordings",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := AlbumsByArtist(db, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := AlbumByID(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := AddAlbum(db, Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}
