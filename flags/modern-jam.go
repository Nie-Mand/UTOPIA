package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetModernJamFlag() string {
	flag := os.Getenv("MODERN_JAM")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}