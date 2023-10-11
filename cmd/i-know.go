package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	I_KNOW_NC = "nc 20.106.177.30 8010"
	I_KNOW_CODE = "https://drive.google.com/file/d/1P81IzSit5U5g6ywQj9lYY2so0XcOSVgl/view?usp=sharing"
	I_KNOW_KEY = "key_for_i_knowwwowowo"
)

func IKnow(
	c* net.Conn,
) {

	statement := `
Tell me what you know about dreaming, tell me have u even had failed?
Tell me do u know freedom, have u ever been lost in a jail!
Im in a red room, my eyes are red too, but got no one to bail!
I gotta evaluate my life, None would like to hear this tale!

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("TCP Stuff -> " + I_KNOW_NC))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("This HOw I did itt -> " + I_KNOW_CODE))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == I_KNOW_KEY {
		(*c).Write([]byte(
			"YOU DOO KNOOOOW -> " + 
			flags.GetIKnowFlag()+ "\n"))
	} else {
		(*c).Write([]byte("U know none..\n"))
		return
	}
}