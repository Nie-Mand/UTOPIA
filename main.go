package main

import (
	"Securinets/UTOPIA/utils"
	"fmt"
)


func main() {
	art := utils.PrintASCIIART()
	fmt.Println(art)
	tcp := utils.TCPServer{URI: ":8000"}
	tcp.Create()
}