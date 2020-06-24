// Package server creates an FTP server
package server

import (
	"fmt"
	"internal/ui"
	"io"
	"log"
	"net"
)

func listen(address, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	reply := ui.RunCommand("welcome")
	if _, err := io.WriteString(c, reply); err != nil {
		log.Print(err)
	}

}
