package main

import (
	"TPC-server/server"
	"log"
)

func main() {
	tpcserver := server.NewServer(":3000")

	tpcserver.GetMensagemChan()

	log.Fatal(tpcserver.Start())

}
