package main
import (
	"flag"
	"log"
    "net"
	"strconv"
	"time"
)

func receiver(conn net.Conn) {
    for{
		reply := make([]byte, 1024)
		_, err := conn.Read(reply)
		if err != nil {
		    return
		}
		log.Println("Answered:", string(reply))
	}
}

func main() {
	strMsg := "Ping"
	port := flag.Int("port", 9999, "Port to listen")
	host := flag.String("host", "127.0.0.1", "IP to listen")
	flag.Parse()

	tcpAddr, err := net.ResolveTCPAddr("tcp", *host+":"+strconv.Itoa(*port))
	if err != nil {
		log.Panicln(err)
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
	go receiver(conn)
	var i int64 = 0;
	for {
		msg := []byte(strMsg+":"+strconv.FormatInt(i, 10))
		_, err = conn.Write(msg)
		if err != nil {
			break
		}
		log.Println("Sended:", string(msg))
		i += 1;
		time.Sleep(100 * time.Millisecond)
	}
}
