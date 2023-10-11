package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	LOOOVE_URL = "https://drive.google.com/file/d/1MXUNjfZ6oxNV6wl_fEp3EH_hFCK58vVQ/view?usp=sharing"
	LOOOVE_URL_2 = "https://drive.google.com/file/d/1b6tALm8yCvTlF7DqGJztcPKYLsclXvnm/view?usp=sharing"
	LOOOVE_KEY = "lithuania"
	LOOOVE_KEY_2 = "alexandra_popp"
)

func Looove(
	c* net.Conn,
) {

	statement := `
Gotta love where I was, Can't leave my homies,
they always with me, can't leave my bros,
I got some memories, I can't compete, with my bruh we bou to flee,
and the sport it is, forever I'll always follow

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("from which song this was taken -> " + LOOOVE_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == LOOOVE_KEY {
		(*c).Write([]byte(
			"Nice, Lithuania is a cool country. Anyways you follow sports -Obv-\n"))
	
			(*c).Write([]byte("Who scored the winner in this match (full name, with underscore)? -> " + LOOOVE_URL_2))
			(*c).Write([]byte("\n"))
			(*c).Write([]byte("\n"))
		(*c).Write([]byte("🎶 ~ key = "))
		key, err := bufio.NewReader(*c).ReadString('\n')
		if err != nil { return }
		key = strings.TrimSpace(string(key))
		
		if key == LOOOVE_KEY_2 {
			(*c).Write([]byte(
				"WE DOOO LOOOVE THAATT, STRAIGHT UP, Here's your flag -> " + 
				flags.GetLooveFlag()+ "\n"))
		} else {
			(*c).Write([]byte("No Loove for a non-fan\n"))
			return
		}

	} else {
		(*c).Write([]byte("No Loove for a non-fan\n"))
		return
	}

}