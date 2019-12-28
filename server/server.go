package server

import (
	"fmt"
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
	data := make(chan []byte, 200000)
	s := NewScreenCapturer(0, data)
	for {
		start := time.Now().UnixNano()
		go s.Capture()
		sendTime := 0
		for {
			frame := <-data
			if frame != nil {
				sendStart := time.Now().UnixNano()
				_, err := c.Write(frame)
				if err != nil {
					//log.Print(err)
					//log.Printf(" frame size : %d bytes, write bytes : %d", len(frame), r)
				}
				sendTime += (int)(time.Now().UnixNano() - sendStart) / 1e3
			} else {
				break
			}
		}
		time.Sleep(1 * time.Millisecond)
		end := (time.Now().UnixNano() - start) / 1e6
		fmt.Printf("FPS: %4.1f, SendBlock: %4d, Capture: %4d (ms), Diff&JPEG: %4d (ms) Send: 0.%4d (ms)\r",
			1000.0/float32(end), s.diff, s.captureTime, s.diffTime, sendTime)
	}

}
