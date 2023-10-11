package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	TOPIA_TWINS_URL = "https://drive.google.com/file/d/1ga3kiYhPLIe6yukdxJOYE35uiDpZVVQF/view?usp=sharing"
	TOPIA_TWINS_KEY = "key_in_utopia_we_got_soldiers"
)
func TopiaTwins(
	c* net.Conn,
) {

	statement := `
Twenty one, Twenty one, Twenty one, straight up..
Im not a hero, im a supervillain, I got more than eight, im on a rolling million,
I got ur key, and it cant go and now im bout to Rock, -you cant rock-,
come save your twin, I got it here hidden, the girl is stuck..

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("download this file -> " + TOPIA_TWINS_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == TOPIA_TWINS_KEY {
		(*c).Write([]byte(
			"THE TWIIINS, STRAIGHT UP, Here's your flag -> " + 
			flags.GetTopiaTwinsFlag()+ "\n"))
	} else {
		(*c).Write([]byte("Your twin won't live to see the Utopia\n"))
		return
	}

}