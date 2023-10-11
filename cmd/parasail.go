package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	PARASAIL_URL = "https://drive.google.com/file/d/10FqFrGuroya6yB2jUH68B5ndGO7GiIsX/view?usp=sharing"
	PARASAIL_KEY = "key_i_s1and_t4ll__I_gEt_UP"
)
func Parasail(
	c* net.Conn,
) {

	statement := `
we gon get there, we gon rock em all,
we can see em all
Guide Him In Da Road, 
Alpha, on the mode
I Do Acknowledge, finally on road.

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("reverese this file -> " + PARASAIL_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == PARASAIL_KEY {
		(*c).Write([]byte(
			"I Forgive Myself, I Forgive Myself -> " + 
			flags.GetParasailFlag()+ "\n"))
	} else {
		(*c).Write([]byte("Wonder what you gonna do..\n"))
		return
	}

}