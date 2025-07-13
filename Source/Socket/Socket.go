package Socket

import (
	"fmt"
	"net"
	"os"
)

func SocketListener() (net.Listener, error) {
	if _, err := os.Stat("/run/applicator.sock"); err == nil {
		err := os.Remove("/run/applicator.sock")
		
		if err != nil {
			return nil, err
		}
	}

	println("Starting socket listener...")
	listener, err := net.Listen("unix", "/run/applicator.sock")

	if err != nil {
		return nil, err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return nil, err
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		return
	}

	data := buffer[:n]

	fmt.Printf("Received data: %s\n", string(data))
	conn.Write([]byte("Data received successfully: " + string(data)))
	conn.Close()
}
