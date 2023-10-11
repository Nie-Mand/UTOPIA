//

package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	THANK_GOD_ENCRYPTED_KEY = "Zmh0Un55Yn9gd1JsUmBkY2J/Um95UnVoflJhZHtkY1JgbGdifw=="
	THANK_GOD_KEY = "key_stormz_a_minor_bt_xes_livin_major"
	THANK_GOD_URL = "https://drive.google.com/file/d/1hsjN4RfyCDT2pVfA-DMVFapP2M71L4kw/view?usp=sharing"
)

func ThankGod(
	c* net.Conn,
) {

	statement := `
Things you hate, but you still you shall find,
if we wont live misery, we'll appreciate the good times,
can't go happy, if I have to drink from the coffee grind,
But I won't deny, I'm still living the good life

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("Reverse it " + THANK_GOD_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("You have this for now " + THANK_GOD_ENCRYPTED_KEY))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == THANK_GOD_KEY {
		(*c).Write([]byte(
			"THAT'S RIGGGHT  -> " + 
			flags.GetThankGodFlag()+ "\n"))
	} else {
		(*c).Write([]byte("You doubted it\n"))
		return
	}
}