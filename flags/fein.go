package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetFeinFlag() string {
	flag := os.Getenv("FEIN")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}