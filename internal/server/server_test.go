package server

import (
	"fmt"
	"net"
	"testing"
)

const testAddress = "localhost"

func TestConnection(t *testing.T) {
	testPort := randPort()
	go listen(testAddress, testPort)
	conn, _ := buildTestConnection(t, testPort)

	if conn != nil {
		defer conn.Close()
	}

}

func TestHandlerReturn(t *testing.T) {
	testPort := randPort()
	go listen(testAddress, testPort)
	conn, _ := buildTestConnection(t, testPort)

	var reply []byte
	b, err := conn.Read(reply)
	if err != nil {
		t.Errorf("unexpected server error: %v", err)
	}
	defer conn.Close()

	t.Log(b)
	t.Log(reply)
}

func buildTestConnection(t *testing.T, testPort string) (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
	if err != nil {
		t.Errorf("server connection failed: %v", err)
	}
	return conn, nil
}

func randPort() string {
	return "8000"
}
