package main

import (
	"log"
	"net"
	"zvon/state"
)

func Listen() {
	ln, err := net.Listen("tcp", "0.0.0.0:8069")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, _ := ln.Accept()
		state.Current.AddConnection(conn)
	}
}
