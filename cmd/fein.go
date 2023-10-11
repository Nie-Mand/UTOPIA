package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	FEIN_URL = "http://20.106.177.30:8008"
	FEIN_CODE_URL = "https://drive.google.com/file/d/1bQXf4Wp7JY3cjLvP8YpkPKyzQ26hMy7l/view?usp=sharing"
	FEIN_KEY = "key_feiiin_feeein_feein_fein_fein"
)

func Fein(
	c* net.Conn,
) {

	statement := `
WHAAAAAAAT, WHAAAAAAT, SLAAT, Uh, uh, uh, yeah,
FEEEIN, KGdasgdas, UYEU, EEEEEEEEE,
MAADE A MILLL, OVER MUMBLINN, AND U CANT KNOW,
WHAATA, MY NEXT SONGG, EEEEEEEEEEEEH,

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("Guess my next song -> " + FEIN_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("This HOw I did itt -> " + FEIN_CODE_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == FEIN_KEY {
		(*c).Write([]byte(
			"WHAAAAAAAAAAAT -> " + 
			flags.GetFeinFlag()+ "\n"))
	} else {
		(*c).Write([]byte("uhuhuh..\n"))
		return
	}
}