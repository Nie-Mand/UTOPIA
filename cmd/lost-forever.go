package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	LOST_FOREVER_NC = "nc 20.106.177.30 8015"
	LOST_FOREVER_CODE = "https://drive.google.com/file/d/1Uuem4hyS8JYsV8SaRBhxBx1O35VVnh_n/view?usp=sharing"
	LOST_FOREVER_KEY = "key_if_u_dont_DH_ull_be_lost"
)

func LostForever(
	c* net.Conn,
) {

	statement := `
Lost in Island, People Violent,
Made the deed, for people who need, but they went tyrant,
I can't speak their lang, cz they dont get my key, 
and if we dont share no key, they'll be always silent...

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("TCP Stuff -> " + LOST_FOREVER_NC))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("This HOw I did itt -> " + LOST_FOREVER_CODE))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))


	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == LOST_FOREVER_KEY {
		(*c).Write([]byte(
			"Ayo!! -> " + 
			flags.GetLostForeverFlag()+ "\n"))
	} else {
		(*c).Write([]byte("I guess we never met..\n"))
		return
	}

}