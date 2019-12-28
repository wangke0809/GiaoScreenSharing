package server

import (
	"bytes"
	"github.com/kbinani/screenshot"
	"image"
	"image/jpeg"
	"time"
)

type Screen struct {
	bounds                                            image.Rectangle
	screenX, screenY, displayIndex, blockStride, diff int
	captureTime, diffTime, flag, quality              int
	buf                                               *bytes.Buffer
	data                                              chan []byte
	lastImg                                           *image.RGBA
}

func NewScreenCapturer(displayIndex, blockSize, quality int, data chan []byte) *Screen {
	s := &Screen{}
	s.bounds = screenshot.GetDisplayBounds(displayIndex)
	s.displayIndex = displayIndex
	s.screenX = s.bounds.Dx()
	s.screenY = s.bounds.Dy()
	s.data = data
	s.buf = new(bytes.Buffer)
	s.lastImg = nil
	s.blockStride = blockSize
	s.quality = quality
	s.flag = 0
	return s
}

func (s *Screen) sendImg(x, y int, i image.Image) {
	s.buf.Reset()
	s.buf.WriteByte(byte(s.blockStride))
	s.buf.WriteByte(byte(s.screenX))
	s.buf.WriteByte(byte(s.screenX >> 8))
	s.buf.WriteByte(byte(s.screenY))
	s.buf.WriteByte(byte(s.screenY >> 8))
	s.buf.WriteByte(byte(x / s.blockStride))
	s.buf.WriteByte(byte(y / s.blockStride))
	o := &jpeg.Options{Quality: 75}
	err := jpeg.Encode(s.buf, i, o)
	if err != nil {
		panic(err)
	}
	s.data <- s.buf.Bytes()
	time.Sleep(10 * time.Microsecond)
}

func (s *Screen) Capture() {
	start := time.Now().UnixNano()
	img, err := screenshot.CaptureRect(s.bounds)
	s.captureTime = (int)(time.Now().UnixNano()-start) / 1e6
	if err != nil {
		panic(err)
	}

	s.flag++
	diff := 0
	start = time.Now().UnixNano()
	for y := 0; y < s.screenY; y += s.blockStride {
		for x := 0; x < s.screenX; x += s.blockStride {
			if x < s.screenX && y < s.screenY {
				endX := Min(x+s.blockStride, s.screenX)
				endY := Min(y+s.blockStride, s.screenY)
				r := image.Rect(x, y, endX, endY)
				subImage := img.SubImage(r)
				isSame := true
				if s.lastImg == nil || s.flag > 100 {
					diff++
					s.sendImg(x, y, subImage)
					continue
				}
				for p := y; p < endY; p += 5 {
					for q := x; q < endX; q += 5 {
						i := img.PixOffset(q, p)
						if img.Pix[i] == s.lastImg.Pix[i] &&
							img.Pix[i+1] == s.lastImg.Pix[i+1] &&
							img.Pix[i+2] == s.lastImg.Pix[i+2] &&
							img.Pix[i+3] == s.lastImg.Pix[i+3] {
						} else {
							isSame = false
							diff++
							s.sendImg(x, y, subImage)
							break
						}
					}
					if !isSame {
						break
					}
				}
			}
		}
	}
	if s.flag > 100 {
		s.flag = 0
	}
	s.diffTime = (int)(time.Now().UnixNano()-start) / 1e6
	s.diff = diff
	s.lastImg = img
	s.data <- nil
}
