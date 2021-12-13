package main

import (
	"kibby/product-cart/database"
	"kibby/product-cart/server"
)

func main() {
	database.Init()
	server.Init()
}
