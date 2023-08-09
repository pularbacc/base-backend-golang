package test

import (
	"testing"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func TestEnvBash(t *testing.T) {
	// fmt.Println("Shell:", os.Getenv("SHELL"))

	fmt.Println("Test Var :", os.Getenv("TEST_VAR"))
}


func TestLibJojoEnv(t *testing.T) {
	// Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    // Get an environment variable
    dbUser := os.Getenv("DB_USER")
    log.Printf("DB user is: %s", dbUser)
}

func TestEnvConst(t *testing.T) {
}