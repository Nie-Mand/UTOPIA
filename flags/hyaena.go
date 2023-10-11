package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetHyaenaFlag() string {
	flag := os.Getenv("HYAENA")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}