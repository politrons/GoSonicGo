package main

import (
	"fmt"
	"image"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

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

var spriteName1 = "enemy-right-1.png"
var spriteName2 = "enemy-right-2.png"
var currentSprite *pixel.Sprite
var count = 0

func run() {

	var i = float64(600)
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

	pic1, err := loadPicture(spriteName1)
	if err != nil {
		panic(err)
	}

	sprite1 := pixel.NewSprite(pic1, pic1.Bounds())

	pic2, err := loadPicture(spriteName2)
	if err != nil {
		panic(err)
	}

	sprite2 := pixel.NewSprite(pic2, pic2.Bounds())

	win.Clear(colornames.Grey)
	//win.MakePicture()

	currentSprite = sprite1

	currentSprite.Draw(win, pixel.IM.Moved(movement))

	for !win.Closed() {
		count = count + 1
		if count%2 == 0 {
			currentSprite = sprite1
		} else {
			currentSprite = sprite2
		}
		i = i + 5
		fmt.Printf("Moving......x:%f y:%f\n", i, i)
		time.Sleep(500 * time.Millisecond)
		win.Update()
		win.Clear(colornames.Grey)
		movement = pixel.Vec{i, 600}
		currentSprite.Draw(win, pixel.IM.Moved(movement))

	}

}

func main() {
	pixelgl.Run(run)
}
