package utils

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type TCPServer struct {
	URI string
}

func (t* TCPServer) HandleConnection(c net.Conn) {
	art := PrintASCIIART()
	c.Write([]byte(art))
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	for {
		c.Write([]byte("Enter Task Code:\n"))
		c.Write([]byte("🎶 ~ "))

		code, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Println(c.RemoteAddr().String(), err)
			return
		}

		code = strings.TrimSpace(string(code))
		bye := RedirectPerCode(code, &c)

		if bye {
			return
		}
	}
}

func (t* TCPServer) Create() {
	l, err := net.Listen("tcp", t.URI)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()


	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go t.HandleConnection(c)
	}
}

