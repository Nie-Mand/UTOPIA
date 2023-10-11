package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetIKnowFlag() string {
	flag := os.Getenv("I_KNOW")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}