package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetSirensFlag() string {
	flag := os.Getenv("SIRENS")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}