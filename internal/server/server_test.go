package server

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net"
	"testing"
	"time"
)

const testAddress = "localhost"

func TestConnection(t *testing.T) {
	testPort := randPort()
	go listen(testAddress, testPort)
	conn, _ := buildTestConnection(t, testPort)
	t.Logf("%v", conn.RemoteAddr())
	defer conn.Close()

}

func TestHandlerReturn(t *testing.T) {
	testPort := randPort()
	time.Sleep(1000 * time.Millisecond)
	go listen(testAddress, testPort)
	conn, _ := buildTestConnection(t, testPort)
	t.Logf("%v", conn.RemoteAddr())
	defer conn.Close()

	reply := new(bytes.Buffer)
	if _, err := io.Copy(reply, conn); err != nil {
		t.Errorf("unexpected server error: %v", err)
	}
	expected := "Welcome!\n"
	if reply.String() != expected {
		t.Errorf("unexpected server reply: want \"%s\", got \"%s\"", expected, reply.String())
	}

}

func buildTestConnection(t *testing.T, testPort string) (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
	if err != nil {
		t.Fatal(err)
	}
	return conn, nil
}

func randPort() string {
	rand.Seed(time.Now().UnixNano())
	for {
		port := rand.Intn(10000)
		if port > 5000 {
			return fmt.Sprintf("%d", port)
		}
	}
}
