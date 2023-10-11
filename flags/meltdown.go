package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetMeltdownFlag() string {
	flag := os.Getenv("MELTDOWN")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}