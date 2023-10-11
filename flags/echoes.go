package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetEchoesFlag() string {
	flag := os.Getenv("DELRESTO")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}