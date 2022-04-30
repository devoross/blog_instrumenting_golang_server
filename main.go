package main

import (
	"instrumented_web_server/server"
)

func main() {
	// start basic web server
	s := server.New(":8080")
	s.Run()
}
