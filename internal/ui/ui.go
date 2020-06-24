// Package ui provides user interface functions for ftp-server
package ui

const ServerWelcome = "Welcome to the Go ftp-server!\n"

// RunCommand runs the input string on the ftp-server
func RunCommand(cmd string) string {
	switch cmd {
	case "welcome":
		return printWelcome() + printHelp()
	case "help":
		return printHelp()
	}
	return ""
}

func printWelcome() string {
	return ServerWelcome
}

func printHelp() string {
	return "Available Commands:\n" +
		"\t'cd' - change directory\n" +
		"\t'ls' - list files\n" +
		"\t'get {filename}' - download file\n" +
		"\t'put {filename}' - upload local file"
}
