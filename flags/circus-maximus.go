package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetCircusMaximusFlag() string {
	flag := os.Getenv("CIRCUS_MAXIMUS")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}