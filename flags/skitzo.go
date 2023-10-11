package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetSkitzoFlag() string {
	flag := os.Getenv("SKITZO")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}