package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetTopiaTwinsFlag() string {
	flag := os.Getenv("TOPIA_TWINS")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}