package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetGodsCountryFlag() string {
	flag := os.Getenv("GODS_COUNTRY")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}