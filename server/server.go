package server

import (
	"log"
	"net"
	"time"
)

func Start(a string) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	data := make(chan []byte)
	s := NewScreenCapturer(0, data)
	for {
		go s.Capture()
		for {
			frame := <- data
			if frame != nil{
				//log.Printf("frame size : %d bytes", len(frame))
				r, err := c.Write(frame)
				r++
				if err != nil {
					log.Fatal(err)
				}
				//log.Printf("r : %d", r)
			}else{
				break
			}
		}

		time.Sleep(30 * time.Millisecond)
		//break
	}

}
