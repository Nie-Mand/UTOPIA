package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetLostForeverFlag() string {
	flag := os.Getenv("LOST_FOREVER")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}