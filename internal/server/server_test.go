package server

import (
	"net"
	"testing"
)

func TestConnection(t *testing.T) {
	go listen()
	conn, _ := buildTestConnection(t)

	if conn != nil {
		defer conn.Close()
	}

}

func buildTestConnection(t *testing.T) (net.Conn, error) {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		t.Errorf("server connection failed: %v", err)
	}
	return conn, nil
}
