package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	MELTDOWN_URL = "https://drive.google.com/file/d/1teR8DuPALSkL4Z2VdZM-WVKi068R54hl/view?usp=sharing"
	MELTDOWN_KEY = "key_drizzy_draaakke_melted_the_chain"
)

func Meltdown(
	c* net.Conn,
) {

	statement := `
I don't need no keys, I got mine in the bucket,
I make feast, I make famine, but your chain ain't worth the ransom,
I walk with cold feet, cold weather, in toronto, with cold sweater,
cold neck, cold wrist and heather, got the keys to of the old boat, cz better,

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("you got those, get the key -> " + MELTDOWN_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == MELTDOWN_KEY {
		(*c).Write([]byte(
			"TORONTO BABYYY -> " + 
			flags.GetMeltdownFlag()+ "\n"))
	} else {
		(*c).Write([]byte("mm, Tensions are definitely rising..\n"))
		return
	}
}