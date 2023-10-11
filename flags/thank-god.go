package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetThankGodFlag() string {
	flag := os.Getenv("THANK_GOD")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}