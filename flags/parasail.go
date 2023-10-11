package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetParasailFlag() string {
	flag := os.Getenv("PARASAIL")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}