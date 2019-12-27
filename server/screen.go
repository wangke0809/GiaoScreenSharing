package server

import (
	"bytes"
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/jpeg"
	"time"
)

type Screen struct {
	bounds                                      image.Rectangle
	screenX, screenY, displayIndex, blockStride int
	buf                                         *bytes.Buffer
	data                                        chan []byte
	lastImg                                     *image.RGBA
}

func NewScreenCapturer(displayIndex int, data chan []byte) *Screen {
	s := &Screen{}
	s.bounds = screenshot.GetDisplayBounds(0)
	s.displayIndex = displayIndex
	s.screenX = s.bounds.Dx()
	s.screenY = s.bounds.Dy()
	s.data = data
	s.buf = new(bytes.Buffer)
	s.lastImg = nil
	s.blockStride = 150
	return s
}

func (s *Screen) Capture() {
	img, err := screenshot.CaptureRect(s.bounds)
	if err != nil {
		panic(err)
	}

	if s.lastImg != nil {
		same := 0
		diff := 0
		for y := 0; y < s.screenY; y += s.blockStride {
			for x := 0; x < s.screenX; x += s.blockStride {
				if x < s.screenX && y < s.screenY {
					endX := Min(x+s.blockStride, s.screenX)
					endY := Min(y+s.blockStride, s.screenY)
					//r := image.Rect(x, y, endX, endY)
					//fmt.Println(x, y, endX, endY, x/s.blockStride, y/s.blockStride, num)
					//subImage := img.SubImage(r)
					//fmt.Println("sub img ", subImage.Bounds().Dx(), subImage.Bounds().Dy())
					//s.buf.Reset()
					//s.buf.WriteByte(byte(num))
					//s.buf.WriteByte(byte(s.screenX))
					//s.buf.WriteByte(byte(s.screenX >> 8))
					//s.buf.WriteByte(byte(s.screenY))
					//s.buf.WriteByte(byte(s.screenY >> 8))
					//s.buf.WriteByte(byte(x / s.blockStride))
					//s.buf.WriteByte(byte(y / s.blockStride))
					//err = jpeg.Encode(s.buf, subImage, nil)
					//if err != nil {
					//	panic(err)
					//}
					//println("buf size :", s.buf.Len())
					//s.data <- s.buf.Bytes()
					//num++
					//time.Sleep(100 * time.Microsecond)

					isSame := true
					for p := y; p < endY; p++ {
						for q := x; q < endX; q++ {
							i := img.PixOffset(q, p)
							if img.Pix[i] == s.lastImg.Pix[i] &&
								img.Pix[i+1] == s.lastImg.Pix[i+1] &&
								img.Pix[i+2] == s.lastImg.Pix[i+2] &&
								img.Pix[i+3] == s.lastImg.Pix[i+3] {
							} else {
								isSame = false
								diff++
								r := image.Rect(x, y, endX, endY)
								subImage := img.SubImage(r)
								s.buf.Reset()
								s.buf.WriteByte(byte(s.screenX))
								s.buf.WriteByte(byte(s.screenX >> 8))
								s.buf.WriteByte(byte(s.screenY))
								s.buf.WriteByte(byte(s.screenY >> 8))
								s.buf.WriteByte(byte(x / s.blockStride))
								s.buf.WriteByte(byte(y / s.blockStride))
								err = jpeg.Encode(s.buf, subImage, nil)
								if err != nil {
									panic(err)
								}
								s.data <- s.buf.Bytes()
								time.Sleep( 50 * time.Microsecond)
								break
							}
						}
						if !isSame {
							break
						}
					}
					if isSame {
						same++
					}
				}
			}
		}
		fmt.Printf("send block : %2d\r", diff)
	}
	s.lastImg = img
	s.data <- nil
}
