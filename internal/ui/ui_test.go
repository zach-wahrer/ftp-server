package ui

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestOnInitialConnectionOutput(t *testing.T) {
	reply := RunCommand("welcome")
	expected := "Welcome!\n" +
		"Available Commands:\n" +
		"\t'cd' - change directory\n" +
		"\t'ls' - list files\n" +
		"\t'get {filename}' - download file\n" +
		"\t'put {filename}' - upload local file"
	if reply != expected {
		t.Errorf("unexpected func reply: want \"%s\", got \"%s\"", expected, reply)
	}
}

func TestPrintHelp(t *testing.T) {
	reply := RunCommand("help")
	expected := "Available Commands:\n" +
		"\t'cd' - change directory\n" +
		"\t'ls' - list files\n" +
		"\t'get {filename}' - download file\n" +
		"\t'put {filename}' - upload local file"
	if reply != expected {
		t.Errorf("unexpected func reply: want \"%s\", got \"%s\"", expected, reply)
	}
}
