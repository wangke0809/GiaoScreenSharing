package client

import (
	"bytes"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
)

type screenLayout struct {
	screen *canvas.Image
	canvas fyne.CanvasObject
	image  image.Image
}

func (s *screenLayout) Layout(o []fyne.CanvasObject, size fyne.Size) {

	s.screen.Resize(size)
	s.screen.Move(fyne.NewPos(0, 0))
}

func (s *screenLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(200, 200)
}

func (s *screenLayout) animate(canvas fyne.Canvas, data chan []byte) {
	go func() {
		for {
			reader := bytes.NewReader(<-data)
			lb, _ := reader.ReadByte()
			blockStride := int(lb)
			lb, _ = reader.ReadByte()
			hb, _ := reader.ReadByte()
			screenX := int(lb) | (int(hb) << 8)
			lb, _ = reader.ReadByte()
			hb, _ = reader.ReadByte()
			screenY := int(lb) | (int(hb) << 8)
			lb, _ = reader.ReadByte()
			hb, _ = reader.ReadByte()
			posX := int(lb)
			posY := int(hb)
			if s.image == nil {
				s.image = image.NewRGBA(image.Rectangle{
					Min: image.Point{0, 0},
					Max: image.Point{screenX, screenY},
				})
			}
			img, err := jpeg.Decode(reader)
			if err != nil {
				log.Println("JPEG Decode err:", err)
				continue
			}
			draw.Draw(s.image.(*image.RGBA), image.Rectangle{
				Min: image.Point{posX * blockStride, posY * blockStride},
				Max: image.Point{posX*blockStride + img.Bounds().Dx(),
					posY*blockStride + img.Bounds().Dy()}},
				img,
				image.Point{0, 0},
				draw.Src)
			s.screen.Image = s.image
			canvas.Refresh(s.screen)
		}
	}()
}

func (s *screenLayout) render() *fyne.Container {
	s.screen = &canvas.Image{}
	container := fyne.NewContainer(s.screen)
	container.Layout = s
	s.canvas = container

	return container
}

func Show(data chan []byte) {
	app := app.New()

	screen := &screenLayout{}
	canvas := screen.render()
	w := app.NewWindow("ScreenSharing")
	go screen.animate(w.Canvas(), data)

	w.SetContent(canvas)

	w.Show()
	w.Canvas().Refresh(screen.canvas)
	app.Run()
}
