package cmd

import (
	"Securinets/UTOPIA/flags"
	"bufio"
	"net"
	"strings"
)

const (
	GODS_COUNTRY_URL = "http://20.106.177.30:8006"
	GODS_COUNTRY_CODE = "https://drive.google.com/file/d/15yZsAEbA7UwfswxQhCGmLmw5_8VS0Y9d/view?usp=sharing"
	GODS_COUNTRY_KEY = "key_in_gods_couuntry_this_is_waaaaaaaar"
)

func GodsCountry(
	c* net.Conn,
) {

	statement := `
LAAA, LAAA LAAA, LAAAA,
People livin like they versions
LAAA, LAAA LAAA, LAAAA,
THey be hiding all their visions
LAAAA, LAAA, LAAAA,
Bout to show em how I can see it
LAAAA, LAAA, LAAAA,
Im bout to hop on on the mission!

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte("SQL Injection and All -> " + GODS_COUNTRY_URL))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("This How I did itt -> " + GODS_COUNTRY_CODE))
	(*c).Write([]byte("\n"))
	(*c).Write([]byte("\n"))

	(*c).Write([]byte("🎶 ~ key = "))
	key, err := bufio.NewReader(*c).ReadString('\n')
	if err != nil { return }
	key = strings.TrimSpace(string(key))

	if key == GODS_COUNTRY_KEY {
		(*c).Write([]byte(
			"WAKIN UP I C THE LIIGHTS -> " + 
			flags.GetGodsCountryFlag()+ "\n"))
	} else {
		(*c).Write([]byte("La, la, la..\n"))
		return
	}
}