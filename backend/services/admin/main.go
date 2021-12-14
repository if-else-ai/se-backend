package main

import (
	"kibby/admin/database"
	"kibby/admin/server"
)

// main
func main() {
	database.Init()
	server.Init()
}
