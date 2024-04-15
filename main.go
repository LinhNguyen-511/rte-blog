package main

import (
	"rte-blog/server"
	"rte-blog/services"
)

func main() {
	services.LoadEnv()
	server.Init()
}
