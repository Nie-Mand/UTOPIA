package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetKPOPFlag() string {
	flag := os.Getenv("KPOP")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}