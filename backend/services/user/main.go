package main

import (
	"kibby/user/database"
	"kibby/user/server"
)

// main
func main() {
	database.Init()
	server.Init()
}
