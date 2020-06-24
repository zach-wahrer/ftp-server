package server

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
	"time"
)

const testAddress = "localhost"
const testPort = "8000"

func TestMain(m *testing.M) {
	go listen(testAddress, testPort)
	time.Sleep(100 * time.Millisecond)
	code := m.Run()
	os.Exit(code)
}

func TestConnectionAndHandlerReturn(t *testing.T) {
	conn, err := createTestConnection(t, testPort)
	if err != nil {
		t.Fatal(err)
	}
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

func createTestConnection(t *testing.T, testPort string) (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprintf("%s:%s", testAddress, testPort))
}
