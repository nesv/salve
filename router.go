package main

import (
	"bytes"
	"log"
	"net"
)

/*
Starts the TCP listener that accepts incoming connections and dispatches them
to the HandleConnection() function so that the requests can be properly routed
across the connected nodes.
*/
func StartRouter(addr string) error {
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting listener on %q (%q)", addr, laddr)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Panicln(err)
		}

		go HandleConnection(conn)
	}
}

/*
Handles any established TCP connections.

This function will also parse any received commands (intended to be sent
directly to a Redis server), create a SHA-1 hash of the provided key so as to
be able to determine which node on the ring has the key.
*/
func HandleConnection(conn *net.TCPConn) {
	buf := new(bytes.Buffer)
	for {
		data := make([]byte, 256)
		n, err := conn.Read(data)
		if err != nil {
			log.Panic(err)
		}

		buf.Write(data)

		if n < 256 {
			// Less than 256 bytes were read, which (hopefully)
			// indicates the message was short enough to fit into
			// a single, buffering pass.
			reply, err := DispatchRequest(buf.Bytes())
		}
	}
	conn.Close()
	return
}

/*
DispatchRequest parses the request enough to get the command name, and the
key name out of the received request.

Once this function has the key name parsed out of the message, it hashes the
key to determine on which node, in the ring, the key lives, and then passes the
request out to the appropriate node, unaltered.

This function returns a byte slice that can be directly written to the
*net.TCPConn the request originated from.
*/
func DispatchRequest(p []byte) (r []byte, err error) {
	return
}
