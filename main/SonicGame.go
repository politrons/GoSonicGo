package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"
	"time"
)

func main() {
	pixelgl.Run(run)
}

func run() {

	var step = float64(300)
	var movement = pixel.Vec{step, step}
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	win.Clear(colornames.Grey)

	greenHillPic, err := loadPicture("green_hill.png")
	if err != nil {
		panic(err)
	}
	greenHillBackground := pixel.NewSprite(greenHillPic, pixel.R(-1200, -800, 1200, 800))
	greenHillBackground.Draw(win, pixel.IM)

	var x = 0
	var pixelRec = float64(36.6)
	var minX = float64(0)
	var maxX = pixelRec

	sonicStopPic, err := loadPicture("sonic_stop.png")
	if err != nil {
		panic(err)
	}

	sonicRightPic, err := loadPicture("sonic_right.png")
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
		win.Clear(colornames.Grey)
		greenHillBackground.Draw(win, pixel.IM)
		if win.Pressed(pixelgl.KeyRight) {
			if x < 9 {
				minX = minX + pixelRec
				maxX = maxX + pixelRec
			} else {
				x = 1
				minX = 0
				maxX = pixelRec
			}
			x = x + 1
			fmt.Printf("Moving......xMin:%f xMax:%f\n", minX, maxX)
			sonicSprite := pixel.NewSprite(sonicRightPic, pixel.R(minX, 0, maxX, 50))

			step = step + 5
			movement = pixel.Vec{step, 220}
			sonicSprite.Draw(win, pixel.IM.Moved(movement))
		} else {
			fmt.Printf("Stoped......xMin:%f xMax:%f\n", minX, maxX)
			movement = pixel.Vec{step, 220}
			sonicSprite := pixel.NewSprite(sonicStopPic, pixel.R(0, 0, pixelRec, 50))
			sonicSprite.Draw(win, pixel.IM.Moved(movement))
		}
		time.Sleep(50 * time.Millisecond)
	}

}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
