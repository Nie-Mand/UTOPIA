package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetLooveFlag() string {
	flag := os.Getenv("LOOOVE")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}