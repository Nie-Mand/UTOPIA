package flags

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetTilFurtherNoticeFlag() string {
	flag := os.Getenv("TIL_FURTHER_NOTICE")
	if flag != "" {
		return flag
	}

	return "Securinets{fake_flag}"
}