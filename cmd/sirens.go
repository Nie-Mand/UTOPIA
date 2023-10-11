package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	SIRENS_URL = "https://drive.google.com/file/d/1_66OZNfbpkObtgSNkyJAr1zV5qPu-sgX/view?usp=sharing"
	SIRENS_KEY = "key_now_i_goot_ur_attenti0n"
)
func Sirens(
	c* net.Conn,
) {

	statement := `
It looks perfect, been a good utopia,
but yet the most of yall, just grinding for their food and all,
im grinding for my pool and all, im grinding for the bun,
im grinding to reach the sun, I hid the key, I hid the gun

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("run this file -> " + SIRENS_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == SIRENS_KEY {
		(*c).Write([]byte(
			"NAAAW I GOOT YOOUR ATTENTIOON -> " + 
			flags.GetSirensFlag()+ "\n"))
	} else {
		(*c).Write([]byte("No attention for you\n"))
		return
	}

}