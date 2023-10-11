package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	MY_EYES_URL = "http://20.106.177.30:8005"
	MY_EYES_KEY = "key_tell_u_demise"
)
func MyEyes(
	c* net.Conn,
) {

	statement := `
BLING BLING, WARMING UP, STILL A PRIME,
I LEFT A PLAIN SCENE, RATHER GO, THAN TO FIND
THE SNAKE DOIN THE DEED, GOTTA RELIEVE, FROM ALL DOJOS
TOO EZ FOR ME, IM GOIN TO LEAVE, TO BE A MOJO!

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("what color is her eyes :p -> " + MY_EYES_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == MY_EYES_KEY {
		(*c).Write([]byte(
			"YEA YEA, THAT SHIT RIIGHT -> " + 
			flags.GetMyEyesFlag()+ "\n"))
	} else {
		(*c).Write([]byte("harder to reach..\n"))
		return
	}

}