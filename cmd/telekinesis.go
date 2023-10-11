package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	TELEKINESIS_URL = "https://drive.google.com/file/d/1pyzkYRQRVPyzqYWiuSZ30NYUT0IKfk9f/view?usp=sharing"
	TELEKINESIS_KEY = "key_to_ur_bash_is_to_look_to_the_future"
)

func Telekinesis(
	c* net.Conn,
) {

	statement := `
I caan't get enough, I can't help it thoo,
I would do it sooo, I even got here for youu,
I'm bout that top liiife, profiles..
I'm and X tho, so no body wouuuld telll..
-yeah, yeah, yeah-
OBTAINED, Im bouta blow flames, cold chains, I gotta move the packs R all,
CONTAINED, the pocket is all drained, I can keep up with this whole thing! YEA YEA

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("get the key, contains two halves -> " + TELEKINESIS_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == TELEKINESIS_KEY {
		(*c).Write([]byte(
			"I CAN SEE THE FUUTURE -> " + 
			flags.GetTelekinesisFlag()+ "\n"))
	} else {
		(*c).Write([]byte("You can't see it..\n"))
		return
	}
}