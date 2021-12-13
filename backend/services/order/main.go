package main

import (
	"kibby/order/database"
	"kibby/order/server"
)

func main() {
	database.Init()
	server.Init()
}
