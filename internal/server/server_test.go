package server

import (
	"fmt"
	"net"
	"testing"
)

const address = "localhost"
const port = "8000"

func TestConnection(t *testing.T) {
	go listen()
	conn, _ := buildTestConnection(t)

	if conn != nil {
		defer conn.Close()
	}

}

func buildTestConnection(t *testing.T) (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		t.Errorf("server connection failed: %v", err)
	}
	return conn, nil
}
