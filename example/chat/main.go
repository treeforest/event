package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/treeforest/event"
	"io"
	"net"
)

const (
	address = "localhost:9090"
)

func server() {
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	dispatcher := event.NewEventDispatcher(struct{}{})
	dispatcher.AddEventListener("msg", func(e event.Event) {
		msg := e.Value().([]byte)
		fmt.Println("> ", string(msg))
	})
	dispatcher.AddEventListener("leave", func(e event.Event) {
		addr := e.Value().(net.Addr)
		fmt.Println("> node leave:", addr.String())
	})

	fmt.Println("server is listening at ", l.Addr().String())
	for {
		var conn net.Conn
		conn, err = l.Accept()
		if err != nil {
			if err == io.EOF {
				continue
			}
			panic(err)
		}

		go func(conn net.Conn) {
			n := 0
			r := bufio.NewReader(conn)
			for {
				buf := make([]byte, 1024)
				n, err = r.Read(buf)
				if err != nil {
					if err == io.EOF {
						continue
					}
					dispatcher.DispatchEvent(event.NewEvent("leave", conn.RemoteAddr()))
					return
				}
				dispatcher.DispatchEvent(event.NewEvent("msg", buf[:n]))
			}
		}(conn)
	}
}

func client() {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("client dial server success, please input message to send")

	for {
		var s string
		_, err = fmt.Scanf("%s\n", &s)
		if err != nil {
			panic(err)
		}
		_, _ = conn.Write([]byte(s))
	}
}

func main() {
	mod := flag.String("mod", "server", "server/client")
	flag.Parse()

	if *mod == "server" {
		server()
	} else {
		client()
	}
}
