package client

import (
	"log"
	"net"
)

const (
	maxDatagramSize = 8192
)

func startListen(a string, data chan []byte)  {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(maxDatagramSize)
	for {
		b := make([]byte, maxDatagramSize)
		n, _, err := l.ReadFromUDP(b)
		if err != nil {
			//log.Println("ReadFromUDP failed:", err)
		}
		//log.Println(n, "bytes read from", src)
		//log.Println(hex.Dump(b[:n]))
		data <- b[:n]
	}
}

func Start(a string) {
	data := make(chan []byte, 200000)
	go startListen(a, data)
	Show(data)
}
