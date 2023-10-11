package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetUtopiaFlag() string {
	flag := os.Getenv("UTOPIA")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}