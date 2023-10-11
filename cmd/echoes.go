package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	ECHOES_URL = "http://20.106.177.30:8009"
	ECHOES_CODE = "https://drive.google.com/file/d/1B5glSLMCg5F18CGvvODBC9akQ2f_f8_s/view?usp=sharing"
	ECHOES_KEY = "f8272161b082e19a6b1a3ff793e0bc5d"
	ECHOES_HASH = "b9a78af4734c33b6fc3b17fbd6eb0011"
)

func DelrestoEchoes(
	c* net.Conn,
) {

	statement := `
-Hush, hush, hush...-
The echoes of the past, the echoes in the present, 
the echoes that we lost, the echoes that we found,
I keep following, with my eyes closed,
I keep looking, but cant get the most,
I MD-Fived em and I, I won't ever let gooo..

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("Using this code -> " + ECHOES_CODE))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("Explore this website -> " + ECHOES_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("Find the file with this hash -> " + ECHOES_HASH))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == ECHOES_KEY {
		(*c).Write([]byte(
			"Only echoes that I wait for, and for that I will give this for you -> " + 
			flags.GetEchoesFlag()+ "\n"))
	} else {
		(*c).Write([]byte("It's the let go.. for you\n"))
		return
	}
}