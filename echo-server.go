package main

import (
	"flag"
	"log"
	"net"
	"strconv"
)

func echo(conn net.Conn) {
	log.Println("Accepted new connection.")
	defer conn.Close()
	defer log.Println("Closed connection.")

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[:size]
		log.Println("Recieved:", string(data))
		conn.Write(data)
	}
}

func main() {
	port := flag.Int("port", 9999, "Port to listen")
	host := flag.String("host", "127.0.0.1", "IP to listen")
	flag.Parse()
	
	socket, err := net.Listen("tcp", *host+":"+strconv.Itoa(*port))
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Listening at '"+*host+":", strconv.Itoa(*port))
	defer socket.Close()

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Panicln(err)
		}
		go echo(conn)
	}
}
