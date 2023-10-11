package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	CIRCUS_MAXIMUS_CODE = "https://drive.google.com/file/d/1rsDG9Y_ExliL3Ag8aVeRMifcgRGkbvnH/view?usp=sharing"
	CIRCUS_MAXIMUS_KEY = "7b98ac209b21860ded9d544d98a5c5f8db7b47b7aef025933e86acf66e6d4614"
)

func CircusMaximus(
	c* net.Conn,
) {

	statement := `
I see all those staaars talking, and I was not awakened,
I know we're a block in life, so who are you, If I'm not mistaken, 
we're all a number, we're all an unnaimed, that's why we are forsaken..

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("What is the hash of the last block for the Second manifacturer -> " + CIRCUS_MAXIMUS_CODE + "\n"))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == CIRCUS_MAXIMUS_KEY {
		(*c).Write([]byte(
			"Unreal whatt is thiss!? -> " + 
			flags.GetCircusMaximusFlag()+ "\n"))
	} else {
		(*c).Write([]byte("Out, whoa-oh\n"))
		return
	}
}