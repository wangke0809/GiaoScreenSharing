package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type FlagStruct struct {
	Host                              string
	Port, Display, BlockSize, Quality int
}

func Start(f FlagStruct) {
	srvAddr := f.Host + ":" + strconv.Itoa(f.Port)
	fmt.Printf("start using %s\r\n", srvAddr)
	addr, err := net.ResolveUDPAddr("udp", srvAddr)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	defer c.Close()
	data := make(chan []byte, 200000)
	s := NewScreenCapturer(f.Display, f.BlockSize, f.Quality, data)
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
				sendTime += (int)(time.Now().UnixNano()-sendStart) / 1e3
			} else {
				break
			}
		}
		time.Sleep(1 * time.Millisecond)
		end := (time.Now().UnixNano() - start) / 1e6
		fmt.Printf("FPS: %3.1f, Send Block: %3d, Capture: %3d (ms), Diff&JPEG Compress: %3d (ms) Send: 0.%3d (ms)\r",
			1000.0/float32(end), s.diff, s.captureTime, s.diffTime, sendTime)
	}

}
