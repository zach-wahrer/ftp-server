package server

import (
	"net"
	"testing"
)

func TestConnection(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		t.Errorf("server connection failed: %v", err)
	}
	if conn != nil {
		defer conn.Close()
	}

}
