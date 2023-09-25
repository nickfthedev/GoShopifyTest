package utils

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Write Test
func TestSendMail(t *testing.T) {
	// Load config
	godotenv.Load(".env")
	// Send Mail
	err := SendEmail(os.Getenv("SMTP_USER"), "This is a Test Mail", "This is a Test Message", "text/plain")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
