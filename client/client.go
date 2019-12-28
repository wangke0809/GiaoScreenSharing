package client

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

const (
	maxDatagramSize = 8192
)

type FlagStruct struct {
	Host string
	Port int
}

func startListen(a string, data chan []byte) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	defer l.Close()
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

func Start(f FlagStruct) {
	srvAddr := f.Host + ":" + strconv.Itoa(f.Port)
	fmt.Printf("start listening %s\r\n", srvAddr)
	data := make(chan []byte, 200000)
	go startListen(srvAddr, data)
	Show(data)
}
