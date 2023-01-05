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

	var i = float64(300)
	var movement = pixel.Vec{i, i}
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
	var step = float64(36.6)
	var minX = float64(0)
	var maxX = step
	sonicPic, err := loadPicture("sonic_1.png")
	if err != nil {
		panic(err)
	}
	for !win.Closed() {
		if x < 9 {
			minX = minX + step
			maxX = maxX + step
		} else {
			x = 1
			minX = 0
			maxX = step
		}
		x = x + 1
		fmt.Printf("Moving......xMin:%f xMax:%f\n", minX, maxX)

		sonicSprite := pixel.NewSprite(sonicPic, pixel.R(minX, 0, maxX, 50))

		i = i + 5
		fmt.Printf("Moving......x:%f y:%f\n", i, i)
		time.Sleep(50 * time.Millisecond)
		win.Update()
		win.Clear(colornames.Grey)
		greenHillBackground.Draw(win, pixel.IM)

		movement = pixel.Vec{i, 220}
		sonicSprite.Draw(win, pixel.IM.Moved(movement))

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
