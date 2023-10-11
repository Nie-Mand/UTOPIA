package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetMyEyesFlag() string {
	flag := os.Getenv("MY_EYES")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}