package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	TIL_FURTHER_NOTICE_URL = "https://drive.google.com/file/d/1_pjZ8imnhd6XP4liHcsNO0i3dEhXZgqm/view?usp=sharing"
	TIL_FURTHER_NOTICE_KEY = "xvjLDLxvvcLfLc|`gvw"
)

func GetKey() string {
	key := TIL_FURTHER_NOTICE_KEY

	real_key := ""
	for i := 0; i < len(key); i++ {
		real_key += string(key[i] ^ 0x13)
	}

	return real_key
}


func TilFurtherNotice(
	c* net.Conn,
) {

	statement := `
I got here, thank god, but how I did it?
this is how I did it, I made it, gotta go to get,
went to the end, put my life, I bid it, but now its for you to get it..
	
`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("Reverse me -> " + TIL_FURTHER_NOTICE_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))


	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == GetKey() {
		(*c).Write([]byte(
			"You found it, Where will you go now? -> " + 
			flags.GetTilFurtherNoticeFlag()+ "\n"))

			(*c).Write([]byte("\n❤️  Thanks for playing UTOPIA, Made By Niemand -> https://github.com/Nie-Mand. XO"))
	} else {
		(*c).Write([]byte("And now it's just a blackout..\n"))
		return
	}

}