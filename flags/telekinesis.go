package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetTelekinesisFlag() string {
	flag := os.Getenv("TELEKINESIS")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}