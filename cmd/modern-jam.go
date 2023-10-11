package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	MODERN_JAM_URL = "https://moderrn-jam.vercel.app"
	MODERN_JAM_KEY = "key_u_gotta_stort_here_2_go_big"
)

func ModernJam(
	c* net.Conn,
) {

	statement := `
You gotta start from here, to go big, to go diamond,
more battles, more wins, but with more losses, you gain power,
too ez, its too ez, but when im on that pro level, a starboy gon be a role model

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("Good Old you need more than google just to find me? -> " + MODERN_JAM_URL + "\n"))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == MODERN_JAM_KEY {
		(*c).Write([]byte(
			"And that boy is not a BRILLIANT STARBOOYY? -> " + 
			flags.GetModernJamFlag()+ "\n"))
	} else {
		(*c).Write([]byte("Out, whoa-oh\n"))
		return
	}
}