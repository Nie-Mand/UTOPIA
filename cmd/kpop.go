package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	KPOP_URL = "https://even-korean.vercel.app"
	KPOP_KEY = "안녕_세계"
)

func KPOP(
	c* net.Conn,
) {

	statement := `
I can trust em, I can trust em, they are all the same,
you cant tell no brothers name, I cant tell naaaadaaa,
and if they tell me they doo the same, I cant tell no liiees,
even if they koreaans, you already knowwww....

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("You know what koreans are famous for, right? -> " + KPOP_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == KPOP_KEY {
		(*c).Write([]byte(
			"Even tho Xe koreaan -> " + 
			flags.GetKPOPFlag()+ "\n"))
	} else {
		(*c).Write([]byte("It's the let go.. for you\n"))
		return
	}
}