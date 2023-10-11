package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	SKITZO_URL = "https://drive.google.com/file/d/1fsAz1fYtclOdOrCpQotH6N0fJxr4x5wJ/view?usp=sharing"
	SKITZO_KEY = "key_bling_blinng_skrr"
)
func Skitzo(
	c* net.Conn,
) {

	statement := `
BLING BLING, WARMING UP, STILL A PRIME,
I LEFT A PLAIN SCENE, RATHER GO, THAN TO FIND
THE SNAKE DOIN THE DEED, GOTTA RELIEVE, FROM ALL DOJOS
TOO EZ FOR ME, IM GOIN TO LEAVE, TO BE A MOJO!

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("download this file -> " + SKITZO_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == SKITZO_KEY {
		(*c).Write([]byte(
			"Yeaaah, Thugger approves -> " + 
			flags.GetSkitzoFlag()+ "\n"))
	} else {
		(*c).Write([]byte("LOCKED\n"))
		return
	}

}