package utils

import (
	"os"
)

const (
	ASCII_ART_PATH = "UTOPIA.txt"
)

func PrintASCIIART() string{
	
	art, err := os.ReadFile(ASCII_ART_PATH)

	if err != nil {
		panic(err)
	}

	return string(art)
}