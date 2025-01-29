package main

import (
	"projeto404/Api/Server"
)

func main() {

	rotas := server.ServerParameters()
	rotas.Run()
}
