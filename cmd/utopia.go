package cmd

import (
	"Securinets/UTOPIA/flags"
	"net"
)

func Utopia(
	c *net.Conn,
) {

	statement := `
Welcome, to UTOPIA, the perfect place to be,
You can find here, all the flags you need,
I will be your guide, to the perfect destination,
For each mission, you shall find a key,
you get the key, you get the flag,
For now, I will gift you, the first one so you can see

`
	(*c).Write([]byte(statement))
	(*c).Write([]byte(flags.GetUtopiaFlag() + "\n"))
}
