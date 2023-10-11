package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	HYAENA_URL = "http://20.106.177.30:8001"
	HYAENA_KEY = "key_nye_west"
)

func Hyaena(
	c* net.Conn,
) {

	statement := `
They be fooling y'all, I never GET,
Kanye know it ain't a bet, KANYE got the agent on set,
Always nauctious, but I still have OPTIONS,
I never let nobody in unless they from the XO-TION!

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("Go get the key -> " + HYAENA_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == HYAENA_KEY {
		(*c).Write([]byte(
			"I FEEL BLESSED -> " + 
			flags.GetHyaenaFlag()+ "\n"))
	} else {
		(*c).Write([]byte("The situation we are in at this time is not a good one..\n"))
		return
	}
}