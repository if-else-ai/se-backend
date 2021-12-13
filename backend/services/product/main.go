package main

import (
	"kibby/product/database"
	"kibby/product/server"
)

func main() {
	database.Init()
	server.Init()
}
