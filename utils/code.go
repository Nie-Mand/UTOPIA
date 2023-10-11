package utils

import (
	"Securinets/UTOPIA/cmd"
	"net"
)

const (
	UTOPIA             = "hello-world"
	HYAENA             = "hyaena"
	THANK_GOD          = "thank-god"
	MODERN_JAM         = "modern-jam"
	MY_EYES            = "my-eyes"
	GODS_COUNTRY       = "gods-country"
	SIRENS             = "sirens"
	MELTDOWN           = "meltdown"
	FEIN               = "fe!n"
	DELRESTO_ECHOES    = "delresto-echoes"
	I_KNOW             = "i-know"
	TOPIA_TWINS        = "topia-twins"
	CIRCUS_MAXIMUS     = "circus-maximus"
	PARASAIL           = "parasail"
	SKITZO             = "skitzo"
	LOST_FOREVER       = "lost-forever"
	LOOOVE             = "looove"
	K_POP              = "kpop"
	TELEKINESIS        = "telekinesis"
	TIL_FURTHER_NOTICE = "til-further-notice"
)

func RedirectPerCode(code string, c *net.Conn) bool {
	if code == UTOPIA {
		cmd.Utopia(c)
	} else if code == HYAENA {
		cmd.Hyaena(c)
	} else if code == TOPIA_TWINS {
		cmd.TopiaTwins(c)
	} else if code == LOOOVE {
		cmd.Looove(c)
	} else if code == MELTDOWN {
		cmd.Meltdown(c)
	} else if code == THANK_GOD {
		cmd.ThankGod(c)
	} else if code == CIRCUS_MAXIMUS {
		cmd.CircusMaximus(c)
	} else if code == PARASAIL {
		cmd.Parasail(c)
	} else if code == TELEKINESIS {
		cmd.Telekinesis(c)
	} else if code == SKITZO {
		cmd.Skitzo(c)
	} else if code == MODERN_JAM {
		cmd.ModernJam(c)
	} else if code == LOST_FOREVER {
		cmd.LostForever(c)
	} else if code == TIL_FURTHER_NOTICE {
		cmd.TilFurtherNotice(c)
	} else if code == K_POP {
		cmd.KPOP(c)
	} else if code == GODS_COUNTRY {
		cmd.GodsCountry(c)
	} else if code == FEIN {
		cmd.Fein(c)
	} else if code == SIRENS {
		cmd.Sirens(c)
	} else if code == MY_EYES {
		cmd.MyEyes(c)
	} else if code == DELRESTO_ECHOES {
		cmd.DelrestoEchoes(c)
	} else {
		(*c).Write([]byte("I didn't get that, me ON no mo\n"))
	}
	return true
}
